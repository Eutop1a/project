package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "helloworld/api/project/v1/user"
	"helloworld/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc     *biz.UserUsecase
	logger *log.Helper
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{
		uc:     uc,
		logger: log.NewHelper(log.DefaultLogger),
	}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	token, err := s.uc.Register(ctx, &biz.RegisterInfo{
		UserID:   new(int64),
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResp{
		StatusCode:  200,
		Msg:         "user: " + req.Username + " register successfully.",
		AccessToken: token,
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	token, err := s.uc.Login(ctx, &biz.RegisterInfo{
		Username: req.Username,
		Password: req.Password,
	})
	resp := &pb.LoginResp{
		StatusCode:  200,
		Msg:         "login successfully.",
		AccessToken: token,
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *UserService) Logout(ctx context.Context, req *pb.LogoutReq) (*emptypb.Empty, error) {
	return nil, s.uc.Logout(ctx, req.UserId)
}
