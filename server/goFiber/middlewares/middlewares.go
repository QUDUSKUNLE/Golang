package middlewares

import "github.com/gofiber/fiber/v2"

func Middleware(context *fiber.Ctx) error {
	return context.SendStatus(404)
}
