package middlewares

import "github.com/gofiber/fiber/v2"

func Logger(context *fiber.Ctx) error {
	return context.SendStatus(404)
}
