package router

import (
	"github.com/gofiber/fiber/v2"
	swa "github.com/arsmn/fiber-swagger/v2"
)


func SwaggerRoute(c *fiber.App) {
	swagger := c.Group("/swagger")

	swagger.Get("*", swa.New())
} 
