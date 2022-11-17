package main

import (
	"auth/pkg/middleware"
	"auth/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.Fiber(app)

	routes.PrivateRoutes(app)
	routes.PublicRoutes(app)

	app.Listen(":5001")
}
