package domain

import (
	"helloworld/internal/data/user/model"
	"time"
)

type AdminUserInfo struct {
	UserID   int64
	Username string
	IsAdmin  bool
	CreateAt time.Time
	UpdateAt time.Time
}

func (u *AdminUserInfo) padding(user *model.User) {
	u.UserID = user.ID
	u.Username = user.Username
	u.IsAdmin = user.IsAdmin
	u.CreateAt = user.CreatedAt
	u.UpdateAt = user.UpdatedAt
}

func NewAdminUserInfo(user *model.User) *UserInfo {
	userInfo := &UserInfo{}
	userInfo.padding(user)
	return userInfo
}

func NewAdminUserInfos(user ...*model.User) []*UserInfo {
	userInfos := make([]*UserInfo, len(user))
	for i, u := range user {
		userInfos[i] = new(UserInfo)
		userInfos[i].padding(u)
	}
	return userInfos
}
