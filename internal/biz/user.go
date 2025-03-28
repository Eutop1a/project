package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"helloworld/internal/data/user/model"
	"helloworld/internal/data/user/query"
	"helloworld/internal/domain"
	"helloworld/internal/pkg/ecode/errdef"
	"helloworld/internal/pkg/encrypt"
	"helloworld/internal/pkg/helper"
	"helloworld/internal/pkg/middlewares/auth"
	"helloworld/internal/pkg/snowflake"
)

type UserRepository interface {
	Register(ctx context.Context, user *model.User) error
	CheckPassword(ctx context.Context, password, hashedPassword string) error
	GetUserInfoByUsername(ctx context.Context, username string) (*domain.UserInfo, error)
}

type UserUseCase struct {
	repo   UserRepository
	jwt    *auth.JWTAuth
	logger *log.Helper
}

func NewUserUseCase(repo UserRepository, jwt *auth.JWTAuth, logger log.Logger) (useCase *UserUseCase, err error) {
	useCase = &UserUseCase{
		repo:   repo,
		jwt:    jwt,
		logger: log.NewHelper(logger),
	}
	return
}

func (uc *UserUseCase) Register(ctx context.Context, registerInfo *domain.UserInfo) (accessToken string, err error) {
	newUser, err := uc.handlerRegisterInfo(ctx, registerInfo)
	if err != nil {
		return "", err
	}
	// create new user's records
	err = uc.repo.Register(ctx, newUser)
	if err != nil {
		return "", helper.HandleError(errdef.ErrCreateUserRecordsFailed, err)
	}
	accessToken, _, err = uc.jwt.CreateToken(newUser.ID, newUser.Username, newUser.IsAdmin)
	return accessToken, err
}

func (uc *UserUseCase) Login(ctx context.Context, user *domain.UserInfo) (accessToken string, err error) {
	userinfo, err := uc.repo.GetUserInfoByUsername(ctx, user.Username)
	if err != nil {
		return "", err
	}
	err = uc.repo.CheckPassword(ctx, user.Password, userinfo.Password)
	if err != nil {
		return "", err
	}
	accessToken, _, err = uc.jwt.CreateToken(userinfo.UserID, userinfo.Username, userinfo.IsAdmin)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (uc *UserUseCase) CheckPassword(ctx context.Context, password, hashedPassword string) error {
	return uc.repo.CheckPassword(ctx, password, hashedPassword)
}

func (uc *UserUseCase) Logout(ctx context.Context, userId string) error {
	return nil
}

func (uc *UserUseCase) handlerRegisterInfo(ctx context.Context, registerInfo *domain.UserInfo) (*model.User, error) {
	userdo := query.Q.User
	count, err := userdo.WithContext(ctx).Where(userdo.Username.Eq(registerInfo.Username)).Count()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.HandleError(errdef.INTERNAL_ERROR, err)
		}
	}
	switch {
	case count > 0: //已有同名用户
		return nil, errdef.ErrUserAlreadyExist
	case len(registerInfo.Password) < 6: //密码长度小于6位
		return nil, errdef.ErrPasswordTooShort
	}
	hashedPassword, err := encrypt.EncPassword(registerInfo.Password)
	if err != nil {
		return nil, errdef.ErrEncryptPasswordFailed
	}
	if registerInfo.UserID <= 0 {
		registerInfo.UserID = snowflake.GetID()
	}
	newUser := &model.User{
		ID:           registerInfo.UserID,
		Username:     registerInfo.Username,
		PasswordHash: hashedPassword,
		IsAdmin:      false,
	}
	return newUser, nil
}
