package middlewares

import (
	"github.com/QUDUSKUNLE/gofiber/src/config"
	"time"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func Logger(context *fiber.Ctx) error {
	return context.SendStatus(404)
}

func Next(context *fiber.Ctx) error {
	fmt.Printf("Date: %s", time.Now())
	return context.Next()
}

func SetContentType(context *fiber.Ctx) error {
	context.Set("Content-Type", "text/plain")
	return context.Next()
}

func AuthReq() func(*fiber.Ctx) error {
	cfg := basicauth.Config{
		Users: map[string]string{
			config.Config("USERNAME"): config.Config("PASSWORD"),
		},
	}
	return basicauth.New(cfg);
}
