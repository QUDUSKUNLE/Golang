package router

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/handlers"
)


func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.FiberHome)

	api := app.Group("/v1")
	api.Post("/products", handlers.CreateProduct)
	api.Get("/products", handlers.GetAllProducts)
	api.Get("products/:id", handlers.GetSingleProduct)
	api.Delete("products/:id", handlers.DeleteProduct)
	api.Patch("products/:id", handlers.UpdateProduct)
}
