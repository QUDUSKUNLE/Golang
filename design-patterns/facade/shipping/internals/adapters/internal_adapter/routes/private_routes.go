package routes

import (
	"github.com/QUDUSKUNLE/shipping/internals/adapters/internal_adapter/handlers"
	"github.com/labstack/echo/v4"
)
func PrivateRoutesAdaptor(private *echo.Group, handler *handlers.HTTPHandler) *echo.Group {
	private.POST("/shipments", handler.PostShipping)
	private.GET("/shipments", handler.GetShippings)

	private.PUT("/pickups", handler.UpdatePickUp)
	private.GET("/pickups", handler.CarrierPickUps)
	private.GET("/pickups/:pick_up_id", handler.GetPickUp)

	private.POST("/addresses", handler.PostAddress)
	private.GET("/addresses", handler.GetAddresses)
	private.GET("/addresses/:address_id", handler.GetAddress)

	private.POST("/packagings", handler.PostPackaging)

	private.POST("/parcels", handler.PostParcel)
	private.GET("/parcels", handler.GetParcel)

	private.GET("/rates", handler.Rates)

	return private
}
