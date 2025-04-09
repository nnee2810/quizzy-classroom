package handler

import (
	"quizzy-classroom/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CreateClassroom(c *fiber.Ctx) error
	FilterClassroomMembers(c *fiber.Ctx) error
	InviteMember(c *fiber.Ctx) error
	FilterInvitations(c *fiber.Ctx) error
	RejectInvitation(c *fiber.Ctx) error
}

type handlerImpl struct {
	CreateClassroomUseCase        usecase.CreateClassroomUseCase
	FilterClassroomMembersUseCase usecase.FilterClassroomMembersUseCase
	InviteMemberUseCase           usecase.InviteMemberUseCase
	FilterInvitationsUseCase      usecase.FilterInvitationsUseCase
	RejectInvitationUseCase       usecase.RejectInvitationUseCase
}

type Inject struct {
	CreateClassroomUseCase        usecase.CreateClassroomUseCase
	FilterClassroomMembersUseCase usecase.FilterClassroomMembersUseCase
	InviteMemberUseCase           usecase.InviteMemberUseCase
	FilterInvitationsUseCase      usecase.FilterInvitationsUseCase
	RejectInvitationUseCase       usecase.RejectInvitationUseCase
}

func New(inject *Inject) Handler {
	return &handlerImpl{
		CreateClassroomUseCase:        inject.CreateClassroomUseCase,
		FilterClassroomMembersUseCase: inject.FilterClassroomMembersUseCase,
		InviteMemberUseCase:           inject.InviteMemberUseCase,
		FilterInvitationsUseCase:      inject.FilterInvitationsUseCase,
		RejectInvitationUseCase:       inject.RejectInvitationUseCase,
	}
}
