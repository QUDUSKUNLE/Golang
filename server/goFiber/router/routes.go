package router

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/middlewares"
	"gofiber/handlers"
)


func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.GetHome)

	api := app.Group("/api", middlewares.AuthReq())
	api.Post("/admin", handlers.GetBody)
	api.Post("/products", handlers.CreateProduct)
	api.Get("/products", handlers.GetAllProducts)
	api.Get("products/:id", handlers.GetSingleProduct)
	api.Delete("products/:id", handlers.DeleteProduct)
}
