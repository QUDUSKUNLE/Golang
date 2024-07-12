package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/QUDUSKUNLE/shipping/internal/adapters/handlers"
)

func PublicRoutesAdaptor(e *echo.Group, handler *handlers.HTTPHandler) *echo.Group {
	e.POST("/delivery", handler.DeliveryProduct)
	e.POST("/reject", handler.RejectProduct)
	e.POST("/users", handler.Register)
	e.POST("/login", handler.Login)
	return e
}
