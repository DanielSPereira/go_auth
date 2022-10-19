package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v2"
)

// @JWTProtected func
// Middleware to protect private routes with jwt handlers
func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtware.Config{
		SigningKey:   []byte(os.Getenv("ACCESS_TOKEN_SECRET")),
		ErrorHandler: handleJwtError,
	}

	return jwtware.New(config)
}

func handleJwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Missing or malformed JWT",
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}
