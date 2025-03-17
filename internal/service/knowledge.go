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
	err = k.uc.CreateKnowledge(ctx, req)
	if err != nil {
		return &common.CommonResp{
			StatusCode: 200,
			Msg:        "failed to create knowledge" + err.Error(),
		}, nil
	}
	return &common.CommonResp{
		StatusCode: 200,
		Msg:        "successfully created knowledge",
	}, nil
}

func (k *KnowledgeService) UpdateKnowledge(ctx context.Context,
	req *knowledgev1.UpdateKnowledgeReq) (rsp *common.CommonResp, err error) {
	err = k.uc.UpdateKnowledge(ctx, req)
	if err != nil {
		return &common.CommonResp{
			StatusCode: 200,
			Msg:        "failed to update knowledge" + err.Error(),
		}, nil
	}
	return &common.CommonResp{
		StatusCode: 200,
		Msg:        "successfully updated knowledge",
	}, nil
}

func (k *KnowledgeService) DeleteKnowledge(ctx context.Context,
	req *knowledgev1.DeleteKnowledgeReq) (rsp *common.CommonResp, err error) {
	err = k.uc.DeleteKnowledge(ctx, req.Id)
	if err != nil {
		return &common.CommonResp{
			StatusCode: 200,
			Msg:        "failed to delete knowledge" + err.Error(),
		}, nil
	}
	return &common.CommonResp{
		StatusCode: 200,
		Msg:        "successfully deleted knowledge",
	}, nil
}

func (k *KnowledgeService) GetKnowledge(ctx context.Context,
	req *knowledgev1.GetKnowledgeReq) (rsp *knowledgev1.GetKnowledgeResp, err error) {
	info, err := k.uc.GetKnowledge(ctx, req.Id)
	if err != nil {
		return &knowledgev1.GetKnowledgeResp{
			Common: &common.CommonResp{
				StatusCode: 200,
				Msg:        "failed to get knowledge" + err.Error(),
			},
			Knowledge: nil,
		}, nil
	}
	return &knowledgev1.GetKnowledgeResp{
		Common: &common.CommonResp{
			StatusCode: 200,
			Msg:        "successfully get knowledge",
		},
		Knowledge: info,
	}, nil
}

func (k *KnowledgeService) ListKnowledge(ctx context.Context,
	req *knowledgev1.ListKnowledgeReq) (rsp *knowledgev1.ListKnowledgeResp, err error) {
	infos, err := k.uc.ListKnowledge(ctx, req.PageSize, req.PageNum)
	if err != nil {
		return &knowledgev1.ListKnowledgeResp{
			Common: &common.CommonResp{
				StatusCode: 200,
				Msg:        "failed to get knowledge list" + err.Error(),
			},
			KnowledgeList: nil,
		}, nil
	}
	return &knowledgev1.ListKnowledgeResp{
		Common: &common.CommonResp{
			StatusCode: 200,
			Msg:        "successfully get knowledge list",
		},
		KnowledgeList: infos,
		TotalCount:    int32(len(infos)),
	}, nil
}
