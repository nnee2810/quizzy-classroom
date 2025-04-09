package handler

import (
	"quizzy-classroom/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CreateClassroom(c *fiber.Ctx) error
	FilterClassroomMembers(c *fiber.Ctx) error
	InviteMember(c *fiber.Ctx) error
}

type handlerImpl struct {
	CreateClassroomUseCase        usecase.CreateClassroomUseCase
	FilterClassroomMembersUseCase usecase.FilterClassroomMembersUseCase
	InviteMemberUseCase           usecase.InviteMemberUseCase
}

type Inject struct {
	CreateClassroomUseCase        usecase.CreateClassroomUseCase
	FilterClassroomMembersUseCase usecase.FilterClassroomMembersUseCase
	InviteMemberUseCase           usecase.InviteMemberUseCase
}

func New(inject *Inject) Handler {
	return &handlerImpl{
		CreateClassroomUseCase:        inject.CreateClassroomUseCase,
		FilterClassroomMembersUseCase: inject.FilterClassroomMembersUseCase,
		InviteMemberUseCase:           inject.InviteMemberUseCase,
	}
}
