package handler

import (
	"github.com/gofiber/fiber/v2"
	"quizzy-classroom/usecase"
)

type Handler interface {
	CreateClassroom(c *fiber.Ctx) error
}

type handlerImpl struct {
	CreateClassroomUseCase usecase.CreateClassroomUseCase
}

type Inject struct {
	CreateClassroomUseCase usecase.CreateClassroomUseCase
}

func New(inject *Inject) Handler {
	return &handlerImpl{
		CreateClassroomUseCase: inject.CreateClassroomUseCase,
	}
}
