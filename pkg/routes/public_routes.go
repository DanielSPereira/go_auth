package routes

import (
	"auth/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	api := a.Group("/api/v1/")

	api.Post("/user/sign/in", controllers.SignIn)
	api.Post("/user/sign/up", controllers.SignUp)
}
