package server

import (
	_ "github.com/go-playground/validator/v10"
	"quizzy-classroom/handler"
	"quizzy-classroom/model"
	"quizzy-classroom/provider"
	"quizzy-classroom/repository"
	"quizzy-classroom/usecase"
	"quizzy-classroom/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
)

func Run(serviceConfig *model.ServiceConfig) {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})
	app.Use(recover.New())

	provider, err := provider.Init(serviceConfig)
	if err != nil {
		logger.Error("failed to init provider", zap.Error(err))
		return
	}
	repository := repository.New(provider)
	handler := handler.New(&handler.Inject{
		CreateClassroomUseCase: usecase.NewCreateClassroomUseCase(repository),
	})

	InitRouter(app, handler)
	util.InitValidate()

	if err := app.Listen(":" + serviceConfig.Port); err != nil {
		logger.Error("failed to start server", zap.Error(err))
		return
	}
}
