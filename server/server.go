package server

import (
	"quizzy-classroom/handler"
	"quizzy-classroom/model"
	"quizzy-classroom/provider"

	"github.com/gofiber/fiber/v2"
	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
)

func Run(serviceConfig *model.ServiceConfig) {
	_, err := provider.Init(serviceConfig)
	if err != nil {
		logger.Error("failed to init provider", zap.Error(err))
		return
	}

	var (
		app     = fiber.New(fiber.Config{})
		handler = handler.New(&handler.HandlerInject{})
	)

	InitRouter(app, handler)

	if err := app.Listen(":" + serviceConfig.Port); err != nil {
		logger.Error("failed to start server", zap.Error(err))
		return
	}
}
