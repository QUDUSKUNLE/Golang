package models

import "github.com/gofiber/fiber/v2"

func GetHome(context *fiber.Ctx) error {
	return context.SendString("Hello, World!")
}

func GetAPI(context *fiber.Ctx) error {
	return context.SendString("API Path: " + context.Params("*"))
}

func GetName(context *fiber.Ctx) error {
	if context.Params("name") != "" {
		return context.SendString("Hello " + context.Params("name"))
	}
	return context.SendString("Where is John?")
}
