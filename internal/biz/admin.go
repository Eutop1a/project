package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/data/dal/model"
	"helloworld/internal/domain"
	"helloworld/internal/pkg/middlewares/auth"
	"strconv"
)

type AdminRepository interface {
	// AddAdmin 添加管理员
	AddAdmin(ctx context.Context, id int64) error
	// DeleteUser 删除用户
	DeleteUser(ctx context.Context, id int64) error
	// GetUserList 获取用户列表
	GetUserList(ctx context.Context, page, size int) ([]*model.User, error)
	// RemoveAdmin 删除管理员
	RemoveAdmin(ctx context.Context, id string) error
	// BatchRemoveAdmin 批量删除管理员
	BatchRemoveAdmin(ctx context.Context, id []string) error
}

type AdminUseCase struct {
	repo   AdminRepository
	jwt    *auth.JWTAuth
	logger *log.Helper
}

func NewAdminUseCase(repo AdminRepository, jwt *auth.JWTAuth, logger log.Logger) (useCase *AdminUseCase, err error) {
	useCase = &AdminUseCase{
		repo:   repo,
		jwt:    jwt,
		logger: log.NewHelper(logger),
	}
	return
}

func (uc *AdminUseCase) AddAdmin(ctx context.Context, id string) (err error) {
	userId, _ := strconv.ParseInt(id, 10, 64)
	err = uc.repo.AddAdmin(ctx, userId)
	return
}

func (uc *AdminUseCase) DeleteUser(ctx context.Context, id string) (err error) {
	userId, _ := strconv.ParseInt(id, 10, 64)
	err = uc.repo.DeleteUser(ctx, userId)
	return
}

func (uc *AdminUseCase) GetUserList(ctx context.Context, page, size int) ([]*domain.AdminUserInfo, error) {
	list, err := uc.repo.GetUserList(ctx, page, size)
	if err != nil {
		return nil, err
	}
	var result []*domain.AdminUserInfo
	for _, item := range list {
		result = append(result, &domain.AdminUserInfo{
			UserID:   item.ID,
			Username: item.Username,
			IsAdmin:  item.IsAdmin,
			CreateAt: item.CreatedAt,
			UpdateAt: item.UpdatedAt,
		})
	}
	return result, nil
}

func (uc *AdminUseCase) RemoveAdmin(ctx context.Context, id string) (err error) {
	err = uc.repo.RemoveAdmin(ctx, id)
	return
}

func (uc *AdminUseCase) BatchRemoveAdmin(ctx context.Context, ids []string) (err error) {
	err = uc.repo.BatchRemoveAdmin(ctx, ids)
	return
}
