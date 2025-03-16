package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	common "helloworld/api/project/v1/common"
	questionv1 "helloworld/api/project/v1/question"
	"helloworld/internal/biz"
)

type QuestionService struct {
	questionv1.UnimplementedQuestionServer
	uc     *biz.QuestionUseCase
	logger *log.Helper
}

func NewQuestionService(uc *biz.QuestionUseCase, logger log.Logger) *QuestionService {
	return &QuestionService{
		uc:     uc,
		logger: log.NewHelper(logger),
	}
}

// CreateQuestionChoice 创建选择题
func (q *QuestionService) CreateQuestionChoice(ctx context.Context,
	in *questionv1.QuestionChoiceInfo) (*common.CommonResp, error) {
	err := q.uc.CreateQuestionChoice(ctx)
	if err != nil {
		return &common.CommonResp{
			StatusCode: 200,
			Msg:        err.Error(),
		}, nil
	}
	return &common.CommonResp{
		StatusCode: 200,
		Msg:        "create question choice success",
	}, nil
}

// CreateQuestionFill 创建填空题
func (q *QuestionService) CreateQuestionFill(ctx context.Context,
	in *questionv1.QuestionFillInfo) (*common.CommonResp, error) {
	return &common.CommonResp{
		StatusCode: 200,
		Msg:        "create question fill success",
	}, nil
}

// CreateQuestionJudge 创建判断题
func (q *QuestionService) CreateQuestionJudge(ctx context.Context,
	in *questionv1.QuestionJudgeInfo) (*common.CommonResp, error) {
	return &common.CommonResp{
		StatusCode: 200,
		Msg:        "create question judge success",
	}, nil
}

// CreateQuestionEssay 创建简答题/大题
func (q *QuestionService) CreateQuestionEssay(ctx context.Context,
	in *questionv1.QuestionEssayInfo) (*common.CommonResp, error) {
	return &common.CommonResp{
		StatusCode: 200,
		Msg:        "create question essay success",
	}, nil
}

// DeleteQuestion 删除题目
func (q *QuestionService) DeleteQuestion(ctx context.Context,
	in *questionv1.DeleteQuestionReq) (*common.CommonResp, error) {
	return &common.CommonResp{
		StatusCode: 200,
		Msg:        "delete question success",
	}, nil
}

// UpdateQuestionChoice 更新选择题
func (q *QuestionService) UpdateQuestionChoice(ctx context.Context,
	in *questionv1.QuestionChoiceInfo) (*common.CommonResp, error) {
	return &common.CommonResp{
		StatusCode: 200,
		Msg:        "update question choice success",
	}, nil
}

// UpdateQuestionFill 更新填空题
func (q *QuestionService) UpdateQuestionFill(ctx context.Context,
	in *questionv1.QuestionFillInfo) (*common.CommonResp, error) {
	return &common.CommonResp{
		StatusCode: 200,
		Msg:        "update question fill success",
	}, nil
}

// UpdateQuestionJudge 更新判断题
func (q *QuestionService) UpdateQuestionJudge(ctx context.Context,
	in *questionv1.QuestionJudgeInfo) (*common.CommonResp, error) {
	return &common.CommonResp{
		StatusCode: 200,
		Msg:        "update question judge success",
	}, nil
}

// UpdateQuestionEssay 更新简答题/大题
func (q *QuestionService) UpdateQuestionEssay(ctx context.Context,
	in *questionv1.QuestionEssayInfo) (*common.CommonResp, error) {
	return &common.CommonResp{
		StatusCode: 200,
		Msg:        "update question essay success",
	}, nil
}

// GetQuestionList 获取题目列表
func (q *QuestionService) GetQuestionList(ctx context.Context,
	in *questionv1.GetQuestionListReq) (*questionv1.GetQuestionListResp, error) {
	return &questionv1.GetQuestionListResp{
		CommonResp: &common.CommonResp{
			StatusCode: 200,
			Msg:        "get question list success",
		},
		Questions: []*questionv1.QuestionInfo{},
		Total:     0,
	}, nil
}
