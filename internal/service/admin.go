package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	adminv1 "helloworld/api/project/v1/admin"
	commonv1 "helloworld/api/project/v1/common"
	pb "helloworld/api/project/v1/user"
	userv1 "helloworld/api/project/v1/user"
	"helloworld/internal/biz"
	"strconv"
)

type AdminService struct {
	pb.UnimplementedUserServer
	uc     *biz.AdminUseCase
	logger *log.Helper
}

func NewAdminService(uc *biz.AdminUseCase) *AdminService {
	return &AdminService{
		uc:     uc,
		logger: log.NewHelper(log.DefaultLogger),
	}
}

func (a *AdminService) AddAdmin(ctx context.Context, req *adminv1.AddAdminReq) (*commonv1.CommonResp, error) {
	err := a.uc.AddAdmin(ctx, req.UserId)
	if err != nil {
		return &commonv1.CommonResp{
			StatusCode: 200,
			Msg:        err.Error(),
		}, nil
	}
	return &commonv1.CommonResp{
		StatusCode: 200,
		Msg:        "successfully add admin " + req.UserId,
	}, nil
}

func (a *AdminService) DeleteUser(ctx context.Context, req *adminv1.DeleteUserReq) (*commonv1.CommonResp, error) {
	err := a.uc.DeleteUser(ctx, req.UserId)
	if err != nil {
		return &commonv1.CommonResp{
			StatusCode: 200,
			Msg:        err.Error(),
		}, nil
	}
	return &commonv1.CommonResp{
		StatusCode: 200,
		Msg:        "successfully delete user " + req.UserId,
	}, nil
}

func (a *AdminService) GetUserList(ctx context.Context, req *adminv1.GetUserListReq) (*adminv1.GetUserListResp, error) {
	userList, err := a.uc.GetUserList(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		return &adminv1.GetUserListResp{
			CommonResp: &commonv1.CommonResp{
				StatusCode: 200,
				Msg:        err.Error(),
			},
		}, nil
	}
	result := make([]*userv1.UserInfo, 0, len(userList))
	for _, user := range userList {
		result = append(result, &userv1.UserInfo{
			UserId:     strconv.FormatInt(user.UserID, 10),
			Username:   user.Username,
			IsAdmin:    user.IsAdmin,
			CreateTime: timestamppb.New(user.CreateAt),
			UpdateTime: timestamppb.New(user.UpdateAt),
		})
	}
	return &adminv1.GetUserListResp{
		CommonResp: &commonv1.CommonResp{
			StatusCode: 200,
			Msg:        "successfully get user list",
		},
		Users: result,
		Total: int32(len(userList)),
	}, nil
}

func (a *AdminService) RemoveAdmin(ctx context.Context, req *adminv1.RemoveAdminReq) (*commonv1.CommonResp, error) {
	err := a.uc.RemoveAdmin(ctx, req.UserId)
	if err != nil {
		return &commonv1.CommonResp{
			StatusCode: 200,
			Msg:        err.Error(),
		}, nil
	}
	return &commonv1.CommonResp{
		StatusCode: 200,
		Msg:        "successfully remove admin " + req.UserId,
	}, nil
}

func (a *AdminService) BatchRemoveAdmin(ctx context.Context, req *adminv1.BatchRemoveAdminReq) (*commonv1.CommonResp, error) {
	err := a.uc.BatchRemoveAdmin(ctx, req.UserIdList)
	if err != nil {
		return &commonv1.CommonResp{
			StatusCode: 200,
			Msg:        err.Error(),
		}, nil
	}
	return &commonv1.CommonResp{
		StatusCode: 200,
		Msg:        "successfully batch remove admin ",
	}, nil
}
