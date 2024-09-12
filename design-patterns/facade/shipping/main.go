package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/QUDUSKUNLE/shipping/docs"
	externalServices "github.com/QUDUSKUNLE/shipping/internals/adapters/external_adapter"
	integration "github.com/QUDUSKUNLE/shipping/internals/adapters/external_adapter"
	"github.com/QUDUSKUNLE/shipping/internals/adapters/internal_adapter/config"
	"github.com/QUDUSKUNLE/shipping/internals/adapters/internal_adapter/handlers"
	validationMiddleware "github.com/QUDUSKUNLE/shipping/internals/adapters/internal_adapter/middleware"
	"github.com/QUDUSKUNLE/shipping/internals/adapters/internal_adapter/repository"
	"github.com/QUDUSKUNLE/shipping/internals/adapters/internal_adapter/routes"
	internalServices "github.com/QUDUSKUNLE/shipping/internals/core/services"
	"github.com/labstack/echo-contrib/echoprometheus"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"golang.org/x/time/rate"
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
	
	e.Use(echoprometheus.NewMiddleware("shipping"))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time}, remote_ip=${remote_ip}, latency=${latency}, method=${method}, uri=${uri}, status=${status}, host=${host}\n",
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 4 << 10,
		LogLevel: 0,
	}))
	// Recover servers when break down
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("ALLOW_ORIGIN")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))
 e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(20))))

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
	e.GET("/metrics", echoprometheus.NewHandler())
	// Start the server on port 8080
	if err := e.Start(fmt.Sprintf(":%s", port)); err != nil && !errors.Is(err, http.ErrServerClosed){
		log.Fatal(err)
	}
}
