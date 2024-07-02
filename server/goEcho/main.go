package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/QUDUSKUNLE/goecho/src/config"
	"github.com/QUDUSKUNLE/goecho/src/middleware"
	"github.com/QUDUSKUNLE/goecho/src/router"
)

func main() {
	// Load environment variable
	if err := config.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get Port number from the loaded .env file
	port := os.Getenv("PORT")
	// Create a new echo instance
	e := echo.New()

	// Register validation
	e = middleware.RegisterValidation(e)

	// Register routes
	e = router.RegisterRoutes(e)

	// Start the server on port 8080
	if err := e.Start(fmt.Sprintf(":%s", port)); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
