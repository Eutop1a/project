package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	knowledgev1 "helloworld/api/project/v1/knowledge"
	"helloworld/internal/domain"
	"helloworld/internal/pkg/middlewares/auth"
	"helloworld/internal/pkg/snowflake"
	"strconv"
)

type KnowledgeRepository interface {
	CreateKnowledge(ctx context.Context, req *domain.KnowledgeInfo) error
	UpdateKnowledge(ctx context.Context, req *domain.KnowledgeInfo) error
	DeleteKnowledge(ctx context.Context, id int64) error
	GetKnowledge(ctx context.Context, id int64) (*domain.KnowledgeInfo, error)
	ListKnowledge(ctx context.Context, pageSize, pageNum int) ([]*domain.KnowledgeInfo, error)
}

type KnowledgeUseCase struct {
	repo   KnowledgeRepository
	jwt    *auth.JWTAuth
	logger *log.Helper
}

func NewKnowledgeUseCase(repo KnowledgeRepository, jwt *auth.JWTAuth, logger log.Logger) *KnowledgeUseCase {
	return &KnowledgeUseCase{
		repo:   repo,
		jwt:    jwt,
		logger: log.NewHelper(logger),
	}
}

func (uc *KnowledgeUseCase) CreateKnowledge(ctx context.Context,
	req *knowledgev1.CreateKnowledgeReq) (err error) {
	pid, _ := strconv.ParseInt(req.Knowledge.ParentId, 10, 64)
	genID := snowflake.GetID()
	if pid == 0 {
		pid = genID
	}
	err = uc.repo.CreateKnowledge(ctx, &domain.KnowledgeInfo{
		Id:          genID,
		Name:        req.Knowledge.Name,
		Description: req.Knowledge.Description,
		CreateTime:  timestamppb.Now(),
		UpdateTime:  timestamppb.Now(),
		ParentId:    genID,
	})
	return err
}

func (uc *KnowledgeUseCase) UpdateKnowledge(ctx context.Context,
	req *knowledgev1.UpdateKnowledgeReq) (err error) {
	id, _ := strconv.ParseInt(req.Id, 10, 64)
	err = uc.repo.UpdateKnowledge(ctx, &domain.KnowledgeInfo{
		Id:          id,
		Name:        req.Knowledge.Name,
		Description: req.Knowledge.Description,
		CreateTime:  timestamppb.Now(),
		UpdateTime:  timestamppb.Now(),
	})
	return err
}

func (uc *KnowledgeUseCase) DeleteKnowledge(ctx context.Context,
	idStr string) (err error) {
	id, _ := strconv.ParseInt(idStr, 10, 64)
	err = uc.repo.DeleteKnowledge(ctx, id)
	return err
}

func (uc *KnowledgeUseCase) GetKnowledge(ctx context.Context,
	idStr string) (*knowledgev1.KnowledgeInfo, error) {
	id, _ := strconv.ParseInt(idStr, 10, 64)
	info, err := uc.repo.GetKnowledge(ctx, id)
	if err != nil {
		return nil, err
	}
	return &knowledgev1.KnowledgeInfo{
		Id:          strconv.FormatInt(info.Id, 10),
		Name:        info.Name,
		Description: info.Description,
		CreateTime:  info.CreateTime,
		UpdateTime:  info.UpdateTime,
	}, nil
}

func (uc *KnowledgeUseCase) ListKnowledge(ctx context.Context,
	pageSize, pageNum int32) ([]*knowledgev1.KnowledgeInfo, error) {
	infos, err := uc.repo.ListKnowledge(ctx, int(pageSize), int(pageNum))
	if err != nil {
		return nil, err
	}
	res := make([]*knowledgev1.KnowledgeInfo, 0)
	for _, info := range infos {
		res = append(res, &knowledgev1.KnowledgeInfo{
			Id:          strconv.FormatInt(info.Id, 10),
			Name:        info.Name,
			Description: info.Description,
			CreateTime:  info.CreateTime,
			UpdateTime:  info.UpdateTime,
		})
	}
	return res, nil
}
