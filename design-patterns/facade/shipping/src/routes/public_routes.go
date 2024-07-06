package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/QUDUSKUNLE/shipping/src/handlers"
)

func PublicRoutesAdaptor(e *echo.Echo) *echo.Echo {
	e.POST("/v1/shipping", handlers.ScheduleProduct)
	e.POST("/v1/pickup", handlers.PickupProduct)
	e.POST("/v1/delivery", handlers.DeliveryProduct)
	e.POST("/v1/reject", handlers.RejectProduct)

	e.POST("/v1/users", handlers.Register)
	e.POST("/v1/login", handlers.Login)
	return e
}
