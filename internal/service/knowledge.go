package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	common "helloworld/api/project/v1/common"
	knowledgev1 "helloworld/api/project/v1/knowledge"
	"helloworld/internal/biz"
)

type KnowledgeService struct {
	knowledgev1.UnimplementedKnowledgeServer
	uc     *biz.KnowledgeUseCase
	logger *log.Helper
}

func NewKnowledgeService(uc *biz.KnowledgeUseCase, logger log.Logger) *KnowledgeService {
	return &KnowledgeService{
		uc:     uc,
		logger: log.NewHelper(logger),
	}
}

func (k *KnowledgeService) CreateKnowledge(ctx context.Context,
	req *knowledgev1.CreateKnowledgeReq) (rsp *common.CommonResp, err error) {
	return &common.CommonResp{
		StatusCode: 200,
		Msg:        "successfully created knowledge",
	}, nil
}

func (k *KnowledgeService) DeleteKnowledge(ctx context.Context,
	req *knowledgev1.DeleteKnowledgeReq) (rsp *common.CommonResp, err error) {
	return &common.CommonResp{
		StatusCode: 200,
		Msg:        "successfully created knowledge",
	}, nil
}

func (k *KnowledgeService) GetKnowledge(ctx context.Context,
	req *knowledgev1.GetKnowledgeReq) (rsp *knowledgev1.GetKnowledgeResp, err error) {
	return &knowledgev1.GetKnowledgeResp{
		Common: &common.CommonResp{
			StatusCode: 200,
			Msg:        "successfully created knowledge",
		},
		Knowledge: &knowledgev1.KnowledgeInfo{},
	}, nil
}

func (k *KnowledgeService) ListKnowledge(ctx context.Context,
	req *knowledgev1.ListKnowledgeReq) (rsp *knowledgev1.ListKnowledgeResp, err error) {
	return &knowledgev1.ListKnowledgeResp{
		Common: &common.CommonResp{
			StatusCode: 200,
			Msg:        "successfully created knowledge",
		},
		KnowledgeList: []*knowledgev1.KnowledgeInfo{},
		TotalCount:    0,
	}, nil
}

func (k *KnowledgeService) UpdateKnowledge(ctx context.Context,
	req *knowledgev1.UpdateKnowledgeReq) (rsp *common.CommonResp, err error) {
	return &common.CommonResp{
		StatusCode: 200,
		Msg:        "successfully update knowledge",
	}, nil
}
