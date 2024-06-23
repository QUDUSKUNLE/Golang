package main

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/models"
	"gofiber/middlewares"
	// "log"
)

func main() {

	// if err := 

	// Fiber framework
	app := fiber.New()
	app.Get("/", models.GetHome)
	app.Get("/:name?", models.GetName)
	app.Get("/api/*", models.GetAPI)
	app.Post("/admin", models.GetBody)
	app.Use(middlewares.Logger)
	app.Listen(":3000")
}
