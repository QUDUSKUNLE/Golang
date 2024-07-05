package router

import (
	"github.com/labstack/echo/v4"
	"github.com/QUDUSKUNLE/shipping/src/handlers"
)

func RoutesAdaptor(e *echo.Echo) *echo.Echo {
	g := e.Group("/v1")
	g.POST("/shipping", handlers.ScheduleProduct)
	g.POST("/pickup", handlers.PickupProduct)
	g.POST("/delivery", handlers.DeliveryProduct)
	g.POST("/reject", handlers.RejectProduct)

	g.POST("/users", handlers.NewUser)
	return e
}
