package domain

import "helloworld/internal/data/dal/model"

type UserInfo struct {
	UserID   int64
	Username string
	Password string
	IsAdmin  bool
}

func (u *UserInfo) padding(user *model.User) {
	u.UserID = user.ID
	u.Username = user.Username
	u.Password = user.PasswordHash
	u.IsAdmin = user.IsAdmin
}

func NewUserInfo(user *model.User) *UserInfo {
	userInfo := &UserInfo{}
	userInfo.padding(user)
	return userInfo
}

func NewUserInfos(user ...*model.User) []*UserInfo {
	userInfos := make([]*UserInfo, len(user))
	for i, u := range user {
		userInfos[i] = new(UserInfo)
		userInfos[i].padding(u)
	}
	return userInfos
}
