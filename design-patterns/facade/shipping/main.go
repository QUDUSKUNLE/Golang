package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/QUDUSKUNLE/shipping/src/config"
	"github.com/QUDUSKUNLE/shipping/src/middleware"
	"github.com/QUDUSKUNLE/shipping/src/router"
)

func init() {
 // Load environment variable
 	if err := config.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Get Port number from the loaded .env file
	port := os.Getenv("PORT")
	// Create a new echo instance
	e := echo.New()

	// Plug echo int validationAdaptor
	e = middleware.ValidationAdaptor(e)

	// Plug echo into routesAdaptor
	e = router.RoutesAdaptor(e)

	// Start the server on port 8080
	if err := e.Start(fmt.Sprintf(":%s", port)); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
