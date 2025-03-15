package ecode

const (
	// Intertal errors
	INTERNAL_ErrInternal = 4001 + iota

	// identity errors
	IDENTITY_LoginFailed
	IDENTITY_UserNameEmpty
	IDENTITY_UserPasswordEmpty
	IDENTITY_UserNotFound
	IDENTITY_UserPasswordError
	IDENTITY_UserAlreadyExist
	IDENTITY_PasswordTooShort
	IDENTITY_PasswordNotMatch
	IDENTITY_SignatureTooLong
	IDENTITY_VerifyCodeExpired
	IDENTITY_VerifyCodeNotMatch
	IDENTITY_EncryptPasswordFailed
	IDENTITY_CreateUserDirFailed
	IDENTITY_CreateUserAvatarFailed
	IDENTITY_ReadAvatarFailed
	IDENTITY_SaveAvatarFailed
	IDENTITY_ErrLogoutFailed
	IDENTITY_ErrCreateUserRecordsFailed

	// auth errors
	AUTH_TokenMissing
	AUTH_TokenInvalid
	AUTH_AccessTokenExpired
	AUTH_RefreshTokenExpired
	AUTH_ParseTokenFailed
	AUTH_CreateTokenFailed
	AUTH_CacheAccessTokenFailed

	// userinfo errors
	USERINFO_ForgetPasswordFailed
	USERINFO_ModifyEmailFailed
	USERINFO_ModifyEmailFailed_Get
	USERINFO_ModifyPasswordFailed
	USERINFO_ModifyUsernameFailed
	USERINFO_GetUserInfoFailed
	USERINFO_ModifySignatureFailed
	IDENTITY_ErrUserLoggedOut

	// parse HTTP request's file errors
	HTTP_FormFileFailed
	HTTP_UploadAvatarFailed

	// admin errors
	ADMIN_UserNotFound
	ADMIN_FailedToGetUser
	ADMIN_FailedToUpdateUserd
)
