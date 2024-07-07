package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/QUDUSKUNLE/shipping/src/handlers"
)

func PublicRoutesAdaptor(e *echo.Echo) *echo.Echo {
	e.POST("/delivery", handlers.DeliveryProduct)
	e.POST("/reject", handlers.RejectProduct)

	e.POST("/users", handlers.Register)
	e.POST("/login", handlers.Login)
	return e
}
