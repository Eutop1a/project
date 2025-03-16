package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/pkg/middlewares/auth"
)

type KnowledgeRepository interface {
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
