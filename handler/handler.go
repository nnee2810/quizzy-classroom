package handler

import (
	"quizzy-classroom/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CreateClassroom(c *fiber.Ctx) error        // Tạo lớp học
	FilterClassroomMembers(c *fiber.Ctx) error // Lọc thành viên của lớp học
	InviteMember(c *fiber.Ctx) error           // Mời thành viên vào lớp học
	FilterInvitations(c *fiber.Ctx) error      // Lọc danh sách lời mời
	RejectInvitation(c *fiber.Ctx) error       // Từ chối lời mời
	AcceptInvitation(c *fiber.Ctx) error       // Chấp nhận lời mời
	FilterInvitedMembers(c *fiber.Ctx) error   // Lọc danh sách thành viên đã được mời
	FilterOwnedClassrooms(c *fiber.Ctx) error  // Lọc danh sách lớp học đang sở hữu
}

type handlerImpl struct {
	CreateClassroomUseCase        usecase.CreateClassroomUseCase
	FilterClassroomMembersUseCase usecase.FilterClassroomMembersUseCase
	InviteMemberUseCase           usecase.InviteMemberUseCase
	FilterInvitationsUseCase      usecase.FilterInvitationsUseCase
	RejectInvitationUseCase       usecase.RejectInvitationUseCase
	AcceptInvitationUseCase       usecase.AcceptInvitationUseCase
	FilterInvitedMembersUseCase   usecase.FilterInvitedMembersUseCase
	FilterOwnedClassroomsUseCase  usecase.FilterOwnedClassroomsUseCase
}

type Inject struct {
	CreateClassroomUseCase        usecase.CreateClassroomUseCase
	FilterClassroomMembersUseCase usecase.FilterClassroomMembersUseCase
	InviteMemberUseCase           usecase.InviteMemberUseCase
	FilterInvitationsUseCase      usecase.FilterInvitationsUseCase
	RejectInvitationUseCase       usecase.RejectInvitationUseCase
	AcceptInvitationUseCase       usecase.AcceptInvitationUseCase
	FilterInvitedMembersUseCase   usecase.FilterInvitedMembersUseCase
	FilterOwnedClassroomsUseCase  usecase.FilterOwnedClassroomsUseCase
}

func New(inject *Inject) Handler {
	return &handlerImpl{
		CreateClassroomUseCase:        inject.CreateClassroomUseCase,
		FilterClassroomMembersUseCase: inject.FilterClassroomMembersUseCase,
		InviteMemberUseCase:           inject.InviteMemberUseCase,
		FilterInvitationsUseCase:      inject.FilterInvitationsUseCase,
		RejectInvitationUseCase:       inject.RejectInvitationUseCase,
		AcceptInvitationUseCase:       inject.AcceptInvitationUseCase,
		FilterInvitedMembersUseCase:   inject.FilterInvitedMembersUseCase,
		FilterOwnedClassroomsUseCase:  inject.FilterOwnedClassroomsUseCase,
	}
}
