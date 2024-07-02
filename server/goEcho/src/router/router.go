package router

import (
	"github.com/labstack/echo/v4"
	"github.com/QUDUSKUNLE/goecho/src/handlers"
)

func RegisterRoutes(e *echo.Echo) *echo.Echo {
	e.GET("/", handlers.Home)
	e.GET("/ping", handlers.Pong)
	e.POST("/register", handlers.Register)
	return e
}
