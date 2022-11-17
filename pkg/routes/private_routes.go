package routes

import (
	"auth/app/controllers"
	"auth/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	api := a.Group("/api/v1/")

	api.Get("/user/current", middleware.JWTProtected(), controllers.CurrentUser)
}
