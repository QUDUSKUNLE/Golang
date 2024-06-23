package main

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/middlewares"
	"gofiber/db"
	"log"
	"gofiber/router"
)

func main() {
	// Connect to database
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	// Fiber framework
	app := fiber.New()
	app.Use(middlewares.Logger)
	router.SetupRoutes(app)
	app.Listen(":3000")
}
