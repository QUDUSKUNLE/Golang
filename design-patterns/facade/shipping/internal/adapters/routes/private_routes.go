package routes

import (
	"github.com/QUDUSKUNLE/shipping/internal/adapters/handlers"
	"github.com/labstack/echo/v4"
)
func PrivateRoutesAdaptor(p *echo.Group, handler *handlers.HTTPHandler) *echo.Group {
	p.POST("/shippings", handler.NewShipping)
	p.GET("/shippings", handler.GetShippings)
	p.PUT("/pickups", handler.UpdatePickUp)
	p.GET("/pickups", handler.CarrierPickUps)
	p.GET("", handler.Restricted)
	return p
}
