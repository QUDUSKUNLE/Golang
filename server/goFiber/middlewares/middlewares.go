package middlewares

import (
	"gofiber/config"


	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func Logger(context *fiber.Ctx) error {
	return context.SendStatus(404)
}

func AuthReq() func(*fiber.Ctx) error {
	cfg := basicauth.Config{
		Users: map[string]string{
			config.Config("USERNAME"): config.Config("PASSWORD"),
		},
	}
	return basicauth.New(cfg);
}
