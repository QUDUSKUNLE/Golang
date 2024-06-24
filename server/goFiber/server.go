package main

import (
	"github.com/gofiber/fiber/v2"
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
	// Add Middlewares
	app.Use(middlewares.Next)
	router.SetupRoutes(app)
	app.Listen(":3000")
}
