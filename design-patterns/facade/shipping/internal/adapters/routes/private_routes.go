package routes

import (
	"github.com/QUDUSKUNLE/shipping/internal/adapters/handlers"
	"github.com/labstack/echo/v4"
)
func PrivateRoutesAdaptor(p *echo.Group) *echo.Group {
	p.POST("/shippings", handlers.NewShipping)
	p.GET("/shippings", handlers.GetShippings)
	p.PUT("/pickups", handlers.UpdatePickUp)
	p.GET("", handlers.Restricted)
	return p
}
