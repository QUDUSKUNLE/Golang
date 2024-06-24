package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/helmet/v2"
	// "github.com/gofiber/fiber/v2/middleware/crsf"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gofiber/middlewares"
	"gofiber/database"
	"gofiber/router"
	"log"
)

func main() {
	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	// Custom server header
	config := fiber.Config{
		ServerHeader: "Fiber Application",
		StrictRouting: true, // Strictly routing the routes
		CaseSensitive: true, // Case sensitivity
	}
	// Fiber framework
	app := fiber.New(config)
	// app.Use(middlewares.Logger)
	app.Use(helmet.New()) // Use helmet middlewares for each route
	// app.Use(crsf.New()) // Use CSRF middleware
	app.Use(limiter.New()) // Use Limiter middleware
	app.Use(logger.New()) // Add logger middleware
	app.Use(middlewares.Next) // Add Time Logging Middlewares
	app.Use("", middlewares.SetContentType)
	router.SetupRoutes(app)
	app.Listen(":3000")
}
