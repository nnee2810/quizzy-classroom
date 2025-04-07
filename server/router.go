package server

import (
	"quizzy-classroom/handler"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App, handler handler.Handler) {
	//app.Get("/sign-in", middleware.JWTMiddleware, handler.SignIn)
}
