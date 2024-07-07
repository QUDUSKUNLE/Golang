package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/QUDUSKUNLE/shipping/src/config"
	echojwt "github.com/labstack/echo-jwt/v4"
	validationMiddleware "github.com/QUDUSKUNLE/shipping/src/middleware"
	"github.com/QUDUSKUNLE/shipping/src/routes"
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

	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) // Recover servers when break down

	// Plug echo int validationAdaptor
	e = validationMiddleware.ValidationAdaptor(e)

	// Plug echo into PublicRoutesAdaptor
	e = routes.PublicRoutesAdaptor(e)

	privateRoutes := e.Group("/v1")
	// Set JWT Configuration
	con := config.JWTConfig(os.Getenv("JWT_SECRET_KEY"))
	privateRoutes.Use(echojwt.WithConfig(con))

	// Plug echo into PrivateRoutesAdaptor
	routes.PrivateRoutesAdaptor(privateRoutes)
	// Start the server on port 8080
	if err := e.Start(fmt.Sprintf(":%s", port)); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
