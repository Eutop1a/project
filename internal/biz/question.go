package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/pkg/middlewares/auth"
)

type QuestionRepository interface {
	CreateQuestionChoice(ctx context.Context) error
	CreateQuestionFill(ctx context.Context) error
	CreateQuestionJudge(ctx context.Context) error
	CreateQuestionEssay(ctx context.Context) error
	DeleteQuestion(ctx context.Context) error
	UpdateQuestionChoice(ctx context.Context) error
	UpdateQuestionFill(ctx context.Context) error
	UpdateQuestionJudge(ctx context.Context) error
	UpdateQuestionEssay(ctx context.Context) error
	GetQuestionList(ctx context.Context) (int64, error)
}

type QuestionUseCase struct {
	repo   QuestionRepository
	jwt    *auth.JWTAuth
	logger *log.Helper
}

func NewQuestionUseCase(repo QuestionRepository, jwt *auth.JWTAuth, logger log.Logger) *QuestionUseCase {
	return &QuestionUseCase{
		repo:   repo,
		jwt:    jwt,
		logger: log.NewHelper(logger),
	}
}

func (uc *QuestionUseCase) CreateQuestionChoice(ctx context.Context) (err error) {
	// TODO: Implement this function
	return err
}

func (uc *QuestionUseCase) CreateQuestionFill(ctx context.Context) (err error) {
	return err
}

func (uc *QuestionUseCase) CreateQuestionJudge(ctx context.Context) (err error) {
	return err
}

func (uc *QuestionUseCase) CreateQuestionEssay(ctx context.Context) (err error) {
	return err
}

func (uc *QuestionUseCase) DeleteQuestion(ctx context.Context) (err error) {
	return err
}

func (uc *QuestionUseCase) UpdateQuestionChoice(ctx context.Context) (err error) {
	return err
}

func (uc *QuestionUseCase) UpdateQuestionFill(ctx context.Context) (err error) {
	return err
}

func (uc *QuestionUseCase) UpdateQuestionJudge(ctx context.Context) (err error) {
	return err
}

func (uc *QuestionUseCase) UpdateQuestionEssay(ctx context.Context) (err error) {
	return err
}

func (uc *QuestionUseCase) GetQuestionList(ctx context.Context) (err error) {
	return err
}
