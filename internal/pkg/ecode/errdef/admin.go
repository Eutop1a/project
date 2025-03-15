package errdef

import (
	kerr "github.com/go-kratos/kratos/v2/errors"
	"helloworld/internal/pkg/ecode"
)

func init() {
	ErrFailedToGetUser = kerr.New(ecode.ADMIN_FailedToGetUser, "找不到的对应用户", "找不到的对应用户")
	ErrFailedToUpdateUser = kerr.New(ecode.ADMIN_FailedToUpdateUserd, "添加该用户为管理员失败", "添加该用户为管理员失败")
}

var (
	ErrFailedToGetUser    *kerr.Error
	ErrFailedToUpdateUser *kerr.Error
)
