package auth

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"helloworld/internal/conf"
	"helloworld/internal/pkg/ecode/errdef"
	"helloworld/internal/pkg/helper"
	"strconv"
	"strings"
	"time"
)

type JWTAuth struct {
	Secret string

	AccessExpireTime time.Duration
	//RefreshExpireTime time.Duration
}

func NewJWTAuth(cfg *conf.JWT) *JWTAuth {
	return &JWTAuth{
		Secret:           cfg.Secret,
		AccessExpireTime: time.Duration(cfg.AccessTokenExpiration) * time.Hour,
		//RefreshExpireTime: time.Duration(cfg.RefreshTokenExpiration) * 24 * time.Hour,
	}
}

const (
	BearerTokenPrefix  = "Bearer"
	AccessTokenHeader  = "Authorization"
	RefreshTokenHeader = "Refresh-Token"

	// TokenIssuer is the default token issuer
	TokenIssuer = "ZEY_HUNTER_ETL"
)

// AccessTokenClaims access token claims, which is used to verify access token
type AccessTokenClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func NewAccessTokenClaims() *AccessTokenClaims {
	return &AccessTokenClaims{}
}
func (t *AccessTokenClaims) isExpired() bool {
	// if token.ExpiresAt before now, the token is expired, return true
	//else, return false
	return t.ExpiresAt.Time.Before(time.Now())
}

// padding pads the access token claims
func (t *AccessTokenClaims) padding(userID int64, username string, isAdmin bool, duration time.Duration) {
	t.UserID = userID
	t.Username = username
	t.IsAdmin = isAdmin
	t.Issuer = TokenIssuer
	t.ExpiresAt = jwt.NewNumericDate(time.Now().Add(duration))

}

// getTokenString returns the token string of access token claims,which is padded
func (t *AccessTokenClaims) getTokenString(secret string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", errdef.ErrCreateTokenFailed
	}
	return
}

//// RefreshTokenClaims refresh token claims, which is used to refresh access token
//type RefreshTokenClaims struct {
//	jwt.RegisteredClaims
//}

//func NewRefreshTokenClaims() *RefreshTokenClaims {
//	return &RefreshTokenClaims{}
//}
//func (t *RefreshTokenClaims) isExpired() bool {
//	return t.ExpiresAt.Time.Before(time.Now())
//}
//
//// padding pads the refresh token claims
//func (t *RefreshTokenClaims) padding(duration time.Duration) {
//	t.Issuer = TokenIssuer
//	t.ExpiresAt = jwt.NewNumericDate(time.Now().Add(duration))
//}
//
//// getTokenString returns the token string of refresh token claims,which is padded
//func (t *RefreshTokenClaims) getTokenString(secret string) (tokenString string, err error) {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
//	tokenString, err = token.SignedString([]byte(secret))
//	if err != nil {
//		return "", helper.HandleError(errdef.ErrCreateTokenFailed, err)
//	}
//	return
//}

// CreateToken create access token and refresh token
func (jwtAuthorizer *JWTAuth) CreateToken(userID int64, username string, isAdmin bool) (accessToken, refreshToken string, err error) {
	// Create access token
	atokenClaims := NewAccessTokenClaims()
	atokenClaims.padding(userID, username, isAdmin, jwtAuthorizer.AccessExpireTime)
	accessToken, err = atokenClaims.getTokenString(jwtAuthorizer.Secret)
	if err != nil {
		return "", "", helper.HandleError(errdef.ErrCreateTokenFailed, err)
	}

	// Create refresh token
	//rtokenClaims := NewRefreshTokenClaims()
	//rtokenClaims.padding(jwtAuthorizer.RefreshExpireTime)
	//refreshToken, err = rtokenClaims.getTokenString(jwtAuthorizer.Secret)
	//if err != nil {
	//	return "", "", helper.HandleError(errdef.ErrCreateTokenFailed, err)
	//}
	return
}

func JwtAuth(secret string, accessTokenExpireTime time.Duration, redisClient *redis.Client) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {

				// Get access token from header and check if it is existed in redis.
				var atoken *jwt.Token
				atoken, err = getToken(tr, secret, AccessTokenHeader)
				if err != nil {
					return nil, err
				}

				// Get refresh token from header
				//var rtoken *jwt.Token
				//rtoken, err = getToken(tr, secret, RefreshTokenHeader)
				if err != nil {
					return nil, err
				}

				aclaims, atok := atoken.Claims.(*AccessTokenClaims)
				//rclaims, rtok := rtoken.Claims.(*RefreshTokenClaims)

				if !atok /* || !rtok*/ {
					return nil, errdef.ErrTokenInvalid
				}

				// Check if access token is existed in redis,
				// because user may log out manually, or the refresh token is expired,
				// which will delete access token in redis.
				_, err2 := redisClient.Get(ctx, strconv.FormatInt(aclaims.UserID, 10)).Result()
				aTokenNotExist := errors.Is(err2, redis.Nil)
				if err2 != nil { // Other error occurs
					if aTokenNotExist {
						return nil, errdef.ErrUserLoggedOut
					}
					return nil, helper.HandleError(errdef.INTERNAL_ERROR, err2)
				}

				// Access token is invalid. Check if access token is expired
				if !atoken.Valid {
					// Check if access token is expired
					if aclaims.isExpired() {
						//if rclaims.isExpired() { // refresh token expired，need login again
						//	return nil, errdef.ErrRefreshTokenExpired
						//} else { // refresh token is not expired, need to refresh access token
						//	// refresh access token
						//	atokenClaims := NewAccessTokenClaims()
						//	atokenClaims.padding(aclaims.UserID, aclaims.Username, aclaims.IsAdmin, accessTokenExpireTime)
						//	accessToken, err := atokenClaims.getTokenString(secret)
						//	if err != nil {
						//		return nil, err
						//	}
						//
						//	// update redis cache:
						//	// ? Set the access token to redis cache with the refresh-token's expiration time.
						//	// ? If we set the access-token's expiration time of access token in redis, user has to log in manually and frequently,
						//	// ? in this case, the refresh-token is meaningless.
						//	redisClient.Set(ctx, strconv.FormatInt(aclaims.UserID, 10), accessToken, rclaims.ExpiresAt.Time.Sub(time.Now()))
						//	return nil, helper.HandleError(errdef.ErrAccessTokenExpired.WithMetadata(map[string]string{
						//		"access_token": accessToken,
						//	}), nil)
						//}
					} else {
						return nil, errdef.ErrTokenInvalid
					}
				}
			}
			return handler(ctx, req)
		}
	}
}

func getToken(tr transport.Transporter, secret, headerString string) (tokenClaims *jwt.Token, err error) {
	tokenString := tr.RequestHeader().Get(headerString) //get token from header,which contains "Bearer "
	auths := strings.SplitN(tokenString, " ", 2)        //get raw token
	if len(auths) != 2 || !strings.EqualFold(auths[0], BearerTokenPrefix) {
		return nil, errdef.ErrTokenMissing
	}
	if headerString == AccessTokenHeader {
		atokenClaims := new(AccessTokenClaims)
		tokenClaims, err = jwt.ParseWithClaims(auths[1], atokenClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil {
			// Skip this error case: token is expired, but it doesn't have other error.
			// IN this case, the error will return nil, because the logic of handling expired token is in middleware.
			if strings.EqualFold("token has invalid claims: token is expired", err.Error()) {
				return tokenClaims, nil
			}
			return nil, helper.HandleError(errdef.ErrParseTokenFailed, err)
		}

	} /*else { //headerString == RefreshTokenHeader
		//rtokenClaims := new(RefreshTokenClaims)
		//tokenClaims, err = jwt.ParseWithClaims(auths[1], rtokenClaims, func(token *jwt.Token) (interface{}, error) {
		//	return []byte(secret), nil
		//})

		if err != nil {
			if strings.EqualFold("token has invalid claims: token is expired", err.Error()) {
				return
			}
			return nil, helper.HandleError(errdef.ErrParseTokenFailed, err)
		}

	}*/
	return
}
