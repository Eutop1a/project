package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
)

type questionRepository struct {
	data   *Data
	logger *log.Helper
}

func NewQuestionRepo(data *Data, logger log.Logger) biz.QuestionRepository {
	return &questionRepository{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

func (q questionRepository) CreateQuestionChoice(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (q questionRepository) CreateQuestionFill(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (q questionRepository) CreateQuestionJudge(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (q questionRepository) CreateQuestionEssay(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (q questionRepository) DeleteQuestion(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (q questionRepository) UpdateQuestionChoice(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (q questionRepository) UpdateQuestionFill(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (q questionRepository) UpdateQuestionJudge(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (q questionRepository) UpdateQuestionEssay(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (q questionRepository) GetQuestionList(ctx context.Context) (int64, error) {
	//TODO implement me
	panic("implement me")
}
