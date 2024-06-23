package handlers

import (
	"github.com/gofiber/fiber/v2"
	"encoding/json"
)

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

func GetBody(context *fiber.Ctx) error {
	body := context.Body()
	var result map[string]interface{}
	if err:= json.Unmarshal(body, &result); err != nil {
		return context.Send(body)
	}
	return context.JSON(result)
}
