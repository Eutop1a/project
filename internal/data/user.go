package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"helloworld/internal/biz"
	"helloworld/internal/data/dal/model"
	"helloworld/internal/data/dal/query"
	"helloworld/internal/domain"
	"helloworld/internal/pkg/ecode/errdef"
	"helloworld/internal/pkg/encrypt"
)

type userRepository struct {
	data   *Data
	logger *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepository {
	return &userRepository{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

func (u *userRepository) Register(ctx context.Context, user *model.User) error {
	var err error
	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err = tx.User.Create(user)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) CheckPassword(ctx context.Context, password, hashedPassword string) error {
	ok := encrypt.DecPassword(password, hashedPassword)
	if !ok {
		return errdef.ErrUserPasswordError
	}
	return nil
}

func (u *userRepository) GetUserInfoByUsername(ctx context.Context, username string) (*domain.UserInfo, error) {
	user := query.User
	userdo := user.WithContext(ctx)
	userInfo, err := userdo.Where(user.Username.Eq(username)).
		Select(user.ID, user.Username, user.IsAdmin, user.PasswordHash).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errdef.ErrUserNotFound
		}
		return nil, err
	}
	return domain.NewUserInfo(userInfo), nil
}
