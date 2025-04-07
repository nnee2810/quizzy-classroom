package server

import (
	"github.com/nnee2810/mimi-core/middleware"
	"quizzy-classroom/handler"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App, handler handler.Handler) {
	classroomGroup := app.Group("/classroom", middleware.JWTMiddleware)
	classroomGroup.Post("/create", handler.CreateClassroom)
}
