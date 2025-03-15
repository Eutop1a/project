package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	commonv1 "helloworld/api/project/v1/common"
	"helloworld/api/project/v1/user"
	"helloworld/internal/biz"
	"helloworld/internal/domain"
)

type UserService struct {
	userv1.UnimplementedUserServer
	uc     *biz.UserUseCase
	logger *log.Helper
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{
		uc:     uc,
		logger: log.NewHelper(log.DefaultLogger),
	}
}

func (s *UserService) Register(ctx context.Context, req *userv1.RegisterReq) (*userv1.RegisterResp, error) {
	token, err := s.uc.Register(ctx, &domain.UserInfo{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return &userv1.RegisterResp{
			CommonResp: &commonv1.CommonResp{
				StatusCode: 200,
				Msg:        err.Error(),
			},
			AccessToken: "",
		}, nil
	}
	return &userv1.RegisterResp{
		CommonResp: &commonv1.CommonResp{
			StatusCode: 200,
			Msg:        "user: " + req.Username + " register successfully.",
		},
		AccessToken: token,
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *userv1.LoginReq) (*userv1.LoginResp, error) {
	token, err := s.uc.Login(ctx, &domain.UserInfo{
		Username: req.Username,
		Password: req.Password,
	})
	resp := &userv1.LoginResp{
		CommonResp: &commonv1.CommonResp{
			StatusCode: 200,
			Msg:        "login successfully.",
		},
		AccessToken: token,
	}
	if err != nil {
		return &userv1.LoginResp{
			CommonResp: &commonv1.CommonResp{
				StatusCode: 200,
				Msg:        err.Error(),
			},
			AccessToken: "",
		}, nil
	}
	return resp, nil
}

func (s *UserService) Logout(ctx context.Context, req *userv1.LogoutReq) (*emptypb.Empty, error) {
	return nil, s.uc.Logout(ctx, req.UserId)
}
