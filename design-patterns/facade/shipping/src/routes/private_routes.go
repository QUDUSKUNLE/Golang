package routes


import (
	"github.com/labstack/echo/v4"
	"github.com/QUDUSKUNLE/shipping/src/handlers"
)

func PrivateRoutesAdaptor(p *echo.Group) *echo.Group {
	p.POST("/shipping", handlers.ScheduleShipping)
	p.GET("", handlers.Restricted)
	return p
}