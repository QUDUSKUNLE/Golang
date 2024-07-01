package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/QUDUSKUNLE/gofiber/src/handlers"
)


func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.FiberHome)

	product := app.Group("/v1/products")
	product.Post("", handlers.CreateProduct)
	product.Get("", handlers.GetAllProducts)
	product.Get("/:id", handlers.GetSingleProduct)
	product.Delete("/:id", handlers.DeleteProduct)
	product.Patch("/:id", handlers.UpdateProduct)

	book := app.Group("v1/books")
	book.Get("", handlers.GetAllBooks)
	book.Post("", handlers.CreateBook)
}
