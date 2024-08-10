package routes

import (
	"github.com/QUDUSKUNLE/shipping/internal/adapters/handlers"
	"github.com/labstack/echo/v4"
)
func PrivateRoutesAdaptor(private *echo.Group, handler *handlers.HTTPHandler) *echo.Group {
	private.POST("/shippings", handler.PostShipping)
	private.GET("/shippings", handler.GetShippings)
	private.PUT("/pickups", handler.UpdatePickUp)
	private.GET("/pickups", handler.CarrierPickUps)
	private.POST("/addresses", handler.PostAddress)
	private.GET("/addresses", handler.GetAddresses)
	private.GET("/addresses/:addressID", handler.GetAddress)
	private.POST("/packaging", handler.PostPackaging)
	private.POST("/parcels", handler.PostParcel)

	private.GET("", handler.Restricted)
	return private
}
