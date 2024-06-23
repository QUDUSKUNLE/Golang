package router

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/middlewares"
	"gofiber/handlers"
)


func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.GetHome)

	api := app.Group("/api", middlewares.AuthReq())
	api.Get("/:name?", handlers.GetName)
	api.Get("/school", handlers.GetAPI)
	api.Post("/admin", handlers.GetBody)
}
