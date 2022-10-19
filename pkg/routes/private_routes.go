package routes

import (
	"auth/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	api := a.Group("/api/v1/")

	api.Get("/user/info", middleware.JWTProtected(), func(c *fiber.Ctx) error { return c.SendString("hello world") })
}
