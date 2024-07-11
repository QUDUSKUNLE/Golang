package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/QUDUSKUNLE/shipping/internal/adapters/handlers"
)

func PublicRoutesAdaptor(e *echo.Group) *echo.Group {
	e.POST("/delivery", handlers.DeliveryProduct)
	e.POST("/reject", handlers.RejectProduct)
	e.POST("/users", handlers.Register)
	e.POST("/login", handlers.Login)
	return e
}
