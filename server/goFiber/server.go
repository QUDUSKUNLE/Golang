package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/QUDUSKUNLE/gofiber/middlewares"
	"github.com/QUDUSKUNLE/gofiber/router"
)

func main() {
	// Custom server header
	config := fiber.Config{
		ServerHeader: "Fiber Application",
		StrictRouting: true, // Strictly routing the routes
		CaseSensitive: true, // Case sensitivity
	}
	// Fiber framework
	app := fiber.New(config)
	app.Use(helmet.New()) // Use helmet middlewares for each route
	// app.Use(crsf.New()) // Use CSRF middleware
	app.Use(limiter.New()) // Use Limiter middleware
	app.Use(logger.New()) // Add logger middleware
	app.Use(middlewares.Next) // Add Time Logging Middlewares
	app.Use("", middlewares.SetContentType)
	// router.SwaggerRoute(app)
	router.SetupRoutes(app)
	app.Listen(":3000")
}
