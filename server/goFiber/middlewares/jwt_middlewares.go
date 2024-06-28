package middlewares

import (
	"os"
	"github.com/gofiber/fiber/v2"

  Middleware "github.com/gofiber/jwt/v2"
)

func JWTProtected() func(*fiber.Ctx) error {
	config := Middleware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey: "jwt",
		ErrorHandler: jwtError,
	}
	return Middleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}
