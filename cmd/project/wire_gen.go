// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
	"helloworld/internal/conf"
	"helloworld/internal/data"
	"helloworld/internal/pkg/middlewares/auth"
	"helloworld/internal/server"
	"helloworld/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, jwt *conf.JWT, logger log.Logger) (*kratos.App, func(), error) {
	mySQLQuestionDB := data.NewMySQLQuestion(confData)
	mySQLUserDB := data.NewMySQLUser(confData)
	client := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(mySQLQuestionDB, mySQLUserDB, client, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepository := data.NewUserRepo(dataData, logger)
	jwtAuth := auth.NewJWTAuth(jwt)
	userUseCase, err := biz.NewUserUseCase(userRepository, jwtAuth, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	userService := service.NewUserService(userUseCase)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	adminRepository := data.NewAdminRepo(dataData, logger)
	adminUseCase, err := biz.NewAdminUseCase(adminRepository, jwtAuth, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	adminService := service.NewAdminService(adminUseCase)
	questionRepository := data.NewQuestionRepo(dataData, logger)
	questionUseCase := biz.NewQuestionUseCase(questionRepository, jwtAuth, logger)
	questionService := service.NewQuestionService(questionUseCase, logger)
	knowledgeRepository := data.NewKnowledgeRepo(dataData, logger)
	knowledgeUseCase := biz.NewKnowledgeUseCase(knowledgeRepository, jwtAuth, logger)
	knowledgeService := service.NewKnowledgeService(knowledgeUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, userService, adminService, questionService, knowledgeService, jwt, client, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
