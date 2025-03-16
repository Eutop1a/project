package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
)

type knowledgeRepository struct {
	data   *Data
	logger *log.Helper
}

func NewKnowledgeRepo(data *Data, logger log.Logger) biz.KnowledgeRepository {
	return &knowledgeRepository{
		data:   data,
		logger: log.NewHelper(logger),
	}
}
