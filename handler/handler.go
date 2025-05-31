package handler

import (
	"quizzy-classroom/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CreateClassroom(c *fiber.Ctx) error            // Tạo lớp học
	FilterClassroomMembers(c *fiber.Ctx) error     // Lọc thành viên của lớp học
	UpdateClassroom(c *fiber.Ctx) error            // Cập nhật thông tin lớp học
	GetClassroomDetail(c *fiber.Ctx) error         // Lấy chi tiết thông tin lớp học
	InviteMember(c *fiber.Ctx) error               // Mời thành viên vào lớp học
	FilterInvitations(c *fiber.Ctx) error          // Lọc danh sách lời mời
	RejectInvitation(c *fiber.Ctx) error           // Từ chối lời mời
	AcceptInvitation(c *fiber.Ctx) error           // Chấp nhận lời mời
	FilterClassroomInvitations(c *fiber.Ctx) error // Lọc danh sách thành viên đã được mời
	FilterJoinedClassrooms(c *fiber.Ctx) error     // Lọc danh sách lớp học đã tham gia
	GetClassroomStat(c *fiber.Ctx) error           // Lấy thống kê về lớp học
	DeleteInvitation(c *fiber.Ctx) error           // Xóa lời mời
}

type handlerImpl struct {
	CreateClassroomUseCase            usecase.CreateClassroomUseCase
	InviteMemberUseCase               usecase.InviteMemberUseCase
	FilterClassroomMembersUseCase     usecase.FilterClassroomMembersUseCase
	FilterClassroomInvitationsUseCase usecase.FilterClassroomInvitationsUseCase
	GetClassroomUseCase               usecase.GetClassroomUseCase
	FilterJoinedClassroomsUseCase     usecase.FilterJoinedClassroomsUseCase
	GetClassroomStatUseCase           usecase.GetClassroomStatUseCase
	UpdateClassroomUseCase            usecase.UpdateClassroomUseCase
	FilterInvitationsUseCase          usecase.FilterInvitationsUseCase
	RejectInvitationUseCase           usecase.RejectInvitationUseCase
	AcceptInvitationUseCase           usecase.AcceptInvitationUseCase
	DeleteInvitationUseCase           usecase.DeleteInvitationUseCase
}

type Inject struct {
	CreateClassroomUseCase            usecase.CreateClassroomUseCase
	InviteMemberUseCase               usecase.InviteMemberUseCase
	FilterClassroomMembersUseCase     usecase.FilterClassroomMembersUseCase
	FilterClassroomInvitationsUseCase usecase.FilterClassroomInvitationsUseCase
	GetClassroomUseCase               usecase.GetClassroomUseCase
	FilterJoinedClassroomsUseCase     usecase.FilterJoinedClassroomsUseCase
	GetClassroomStatUseCase           usecase.GetClassroomStatUseCase
	UpdateClassroomUseCase            usecase.UpdateClassroomUseCase
	FilterInvitationsUseCase          usecase.FilterInvitationsUseCase
	RejectInvitationUseCase           usecase.RejectInvitationUseCase
	AcceptInvitationUseCase           usecase.AcceptInvitationUseCase
	DeleteInvitationUseCase           usecase.DeleteInvitationUseCase
}

func New(inject *Inject) Handler {
	return &handlerImpl{
		CreateClassroomUseCase:            inject.CreateClassroomUseCase,
		InviteMemberUseCase:               inject.InviteMemberUseCase,
		FilterClassroomMembersUseCase:     inject.FilterClassroomMembersUseCase,
		FilterClassroomInvitationsUseCase: inject.FilterClassroomInvitationsUseCase,
		GetClassroomUseCase:               inject.GetClassroomUseCase,
		FilterJoinedClassroomsUseCase:     inject.FilterJoinedClassroomsUseCase,
		GetClassroomStatUseCase:           inject.GetClassroomStatUseCase,
		UpdateClassroomUseCase:            inject.UpdateClassroomUseCase,
		FilterInvitationsUseCase:          inject.FilterInvitationsUseCase,
		RejectInvitationUseCase:           inject.RejectInvitationUseCase,
		AcceptInvitationUseCase:           inject.AcceptInvitationUseCase,
		DeleteInvitationUseCase:           inject.DeleteInvitationUseCase,
	}
}
