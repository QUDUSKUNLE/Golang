package main

import "github.com/gofiber/fiber/v2"
import "gofiber/models"

func main() {
	app := fiber.New()

	app.Get("/", models.GetHome)
	app.Get("/:name?", models.GetName)
	app.Get("/api/*", models.GetAPI)
	app.Listen(":3000")
}
