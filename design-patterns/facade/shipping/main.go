package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/QUDUSKUNLE/shipping/internal/adapters/config"
	"github.com/QUDUSKUNLE/shipping/internal/adapters/handlers"
	validationMiddleware "github.com/QUDUSKUNLE/shipping/internal/adapters/middleware"
	"github.com/QUDUSKUNLE/shipping/internal/adapters/repository"
	"github.com/QUDUSKUNLE/shipping/internal/adapters/integration"
	"github.com/QUDUSKUNLE/shipping/internal/adapters/routes"
	"github.com/QUDUSKUNLE/shipping/internal/core/services"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	svc *services.InternalServicesHandler
	ext *services.ExternalServicesHandler
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
	e = validationMiddleware.ValidationAdaptor(e)
	
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) // Recover servers when break down


	store, err := repository.OpenDBConnection()

	if err != nil {
		log.Fatalf("Error connecting to the databse: %s", err.Error())
	}
	extStore := integration.OpenExternalConnection()
	svc = services.InternalServicesAdapter(store)
	ext = services.ExternalServicesAdapter(extStore)
	httpHandler := handlers.HttpAdapter(*svc, *ext)

	// Plug echo into PublicRoutesAdaptor
	public := e.Group("/v1")
	routes.PublicRoutesAdaptor(public, httpHandler)

	privateRoutes := e.Group("/v1")
	// Set JWT Configuration
	con := config.JWTConfig(os.Getenv("JWT_SECRET_KEY"))
	privateRoutes.Use(echojwt.WithConfig(con))

	// Plug echo into PrivateRoutesAdaptor
	routes.PrivateRoutesAdaptor(privateRoutes, httpHandler)
	// Start the server on port 8080
	if err := e.Start(fmt.Sprintf(":%s", port)); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
