package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"helloworld/internal/biz"
	"helloworld/internal/data/user/model"
	"helloworld/internal/data/user/query"
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
	err = tx.Commit()
	return err
}

func (a *adminRepository) DeleteUser(ctx context.Context, id int64) (err error) {
	// 直接执行软删除操作，无需开启事务
	u := query.User
	_, err = u.WithContext(ctx).Where(u.ID.Eq(id)).Delete()
	return err
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
	list, _, err = tx.User.WithContext(ctx).
		Select(query.User.ID, query.User.Username, query.User.IsAdmin, query.User.CreatedAt, query.User.UpdatedAt).
		Where(query.User.DeletedAt.IsNull()).
		Order(query.User.CreatedAt.Desc()).
		FindByPage((page-1)*size, size)
	err = tx.Commit()
	return
}

func (a *adminRepository) RemoveAdmin(ctx context.Context, id int64) (err error) {
	u := query.User
	userDo := u.WithContext(ctx)
	_, err = userDo.Where(u.ID.Eq(id)).UpdateSimple(u.IsAdmin.Value(false))
	return err
}

func (a *adminRepository) BatchRemoveAdmin(ctx context.Context, ids []int64) (err error) {
	// 开启事务
	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		if p := recover(); p != nil {
			tx.Rollback()
		}
	}()
	// 使用批量更新
	_, err = tx.User.WithContext(ctx).
		Where(tx.User.ID.In(ids...)).
		UpdateSimple(tx.User.IsAdmin.Value(false))
	if err != nil {
		return errdef.ErrFailedToUpdateUser
	}
	return tx.Commit()
}
