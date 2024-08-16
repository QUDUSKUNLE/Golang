package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/QUDUSKUNLE/shipping/internals/adapters/internal_adapter/config"
	"github.com/QUDUSKUNLE/shipping/internals/adapters/internal_adapter/handlers"
	validationMiddleware "github.com/QUDUSKUNLE/shipping/internals/adapters/internal_adapter/middleware"
	"github.com/QUDUSKUNLE/shipping/internals/adapters/internal_adapter/repository"
	integration "github.com/QUDUSKUNLE/shipping/internals/adapters/external_adapter/integration/terminals"
	"github.com/QUDUSKUNLE/shipping/internals/adapters/internal_adapter/routes"
	internalServices "github.com/QUDUSKUNLE/shipping/internals/core/services"
	externalServices "github.com/QUDUSKUNLE/shipping/internals/adapters/external_adapter/integration"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/QUDUSKUNLE/shipping/docs"
)

var (
	internal *internalServices.InternalServicesHandler
	external *externalServices.ExternalServicesHandler
)

func init() {
 // Load environment variable
 	if err := config.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @title Bahsoon API
// @version 1.0
// @description Bahsoon API
// @host localhost:8080
// @BasePath /v1
func main() {
	// Get Port number from the loaded .env file
	port := os.Getenv("PORT")
	// Create a new echo instance
	e := echo.New()
	// Plug echo int validationAdaptor
	e = validationMiddleware.ValidationAdaptor(e)
	
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) // Recover servers when break down
	e.Use(middleware.CORS())


	internalStore, err := repository.OpenDBConnection()
	if err != nil {
		log.Fatalf("Error connecting to the databse: %s", err.Error())
	}

	externalStore := integration.OpenExternalConnection()
	internal = internalServices.InternalServicesAdapter(internalStore)
	external = externalServices.ExternalServicesAdapter(externalStore)
	httpHandler := handlers.HttpAdapter(*internal, *external)

	// Plug echo into PublicRoutesAdaptor
	public := e.Group("/v1")
	routes.PublicRoutesAdaptor(public, httpHandler)

	privateRoutes := e.Group("/v1")
	// Set JWT Configuration
	con := config.JWTConfig(os.Getenv("JWT_SECRET_KEY"))
	privateRoutes.Use(echojwt.WithConfig(con))

	// Plug echo into PrivateRoutesAdaptor
	routes.PrivateRoutesAdaptor(privateRoutes, httpHandler)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// Start the server on port 8080
	if err := e.Start(fmt.Sprintf(":%s", port)); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
