package server

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
}
