package http

import (
	"quizzy-classroom/handler"
	"quizzy-classroom/model"
	"quizzy-classroom/provider"
	"quizzy-classroom/repository"
	"quizzy-classroom/usecase"

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
		CreateClassroomUseCase:            usecase.NewCreateClassroomUseCase(repository),
		FilterClassroomMembersUseCase:     usecase.NewFilterClassroomMembersUseCase(repository),
		UpdateClassroomUseCase:            usecase.NewUpdateClassroomUseCase(repository),
		GetClassroomUseCase:               usecase.NewGetClassroomUseCase(repository),
		InviteMemberUseCase:               usecase.NewInviteMemberUseCase(repository),
		FilterInvitationsUseCase:          usecase.NewFilterInvitationsUseCase(repository),
		RejectInvitationUseCase:           usecase.NewRejectInvitationUseCase(repository),
		AcceptInvitationUseCase:           usecase.NewAcceptInvitationUseCase(repository),
		FilterClassroomInvitationsUseCase: usecase.NewFilterClassroomInvitationsUseCase(repository),
		FilterJoinedClassroomsUseCase:     usecase.NewFilterJoinedClassroomsUseCase(repository),
		GetClassroomStatUseCase:           usecase.NewGetClassroomStatUseCase(repository),
		DeleteInvitationUseCase:           usecase.NewDeleteInvitationUseCase(repository),
	})

	InitRouter(app, handler)

	if err := app.Listen(":" + serviceConfig.Port); err != nil {
		logger.Error("failed to start httpr", zap.Error(err))
		return
	}
}
