package grpc

import (
	"github.com/nnee2810/mimi-core/grpc"
	"github.com/nnee2810/mimi-core/grpc/proto/gen"
	grpc2 "google.golang.org/grpc"
	"quizzy-classroom/grpc/handler"
	"quizzy-classroom/model"
	"quizzy-classroom/provider"
	"quizzy-classroom/repository"

	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
)

func Run(serviceConfig *model.ServiceConfig) {
	server := grpc.New()

	provider, err := provider.Init(serviceConfig)
	if err != nil {
		logger.Error("failed to init provider", zap.Error(err))
		return
	}
	repository := repository.New(provider)
	handler := handler.New(repository)

	server.RegisterService(func(server *grpc2.Server) {
		gen.RegisterClassroomServer(server, handler)
	})

	if err := server.Start(":" + serviceConfig.GrpcPort); err == nil {
		logger.Info("grpc server started", zap.String("port", serviceConfig.GrpcPort))
	} else {
		logger.Error("failed to start grpc server", zap.Error(err))
	}
}
