package http

import (
	"quizzy-classroom/handler"

	"github.com/nnee2810/mimi-core/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App, handler handler.Handler) {
	classroomGroup := app.Group("/classroom", middleware.JWTMiddleware)
	classroomGroup.Post("/create", handler.CreateClassroom)
	classroomGroup.Post("/invite", handler.InviteMember)
	classroomGroup.Get("/:classroom_id/members", handler.FilterClassroomMembers)
	classroomGroup.Get("/:classroom_id/invitations", handler.FilterClassroomInvitations)
	classroomGroup.Get("/:classroom_id/detail", handler.GetClassroomDetail)
	classroomGroup.Get("/joined", handler.FilterJoinedClassrooms)
	classroomGroup.Get("/stat", handler.GetClassroomStat)
	classroomGroup.Patch("/update", handler.UpdateClassroom)

	invitationGroup := app.Group("/invitation", middleware.JWTMiddleware)
	invitationGroup.Get("/filter", handler.FilterInvitations)
	invitationGroup.Put("/reject/:id", handler.RejectInvitation)
	invitationGroup.Put("/accept/:id", handler.AcceptInvitation)
	invitationGroup.Delete("/delete/:id", handler.DeleteInvitation)
}
