package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"helloworld/internal/biz"
	"helloworld/internal/data/dal/model"
	"helloworld/internal/data/dal/query"
	"helloworld/internal/pkg/ecode/errdef"
)

type adminRepository struct {
	data   *Data
	logger *log.Helper
}

func NewAdminRepo(data *Data, logger log.Logger) biz.AdminRepository {
	return &adminRepository{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

func (a *adminRepository) AddAdmin(ctx context.Context, userID int64) (err error) {
	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	user, err := tx.User.WithContext(ctx).
		Where(query.User.ID.Eq(userID)).
		Where(query.User.DeletedAt.IsNull()).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errdef.ErrUserNotFound
		}
		return errdef.ErrFailedToGetUser
	}
	if user.IsAdmin {
		return nil
	}
	_, err = tx.User.WithContext(ctx).
		Where(query.User.ID.Eq(userID)).
		UpdateSimple(query.User.IsAdmin.Value(true))
	if err != nil {
		return errdef.ErrFailedToUpdateUser
	}
	return nil

}

func (a *adminRepository) DeleteUser(ctx context.Context, id int64) (err error) {
	return nil
}

func (a *adminRepository) GetUserList(ctx context.Context, page, size int) (list []*model.User, err error) {
	if size <= 0 {
		size = 10 // 默认值或返回错误
	}
	if page <= 0 {
		page = 1
	}
	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	list, _, err = tx.User.WithContext(ctx).Debug().
		Select(query.User.ID, query.User.Username, query.User.IsAdmin, query.User.CreatedAt, query.User.UpdatedAt).
		Where(query.User.DeletedAt.IsNull()).
		Order(query.User.CreatedAt.Desc()).
		FindByPage((page-1)*size, size)
	err = tx.Commit()
	return
}

func (a *adminRepository) RemoveAdmin(ctx context.Context, id string) error {
	return nil
}

func (a *adminRepository) BatchRemoveAdmin(ctx context.Context, id []string) error {
	return nil
}
