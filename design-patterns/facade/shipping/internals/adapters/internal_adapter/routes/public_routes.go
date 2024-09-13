package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/QUDUSKUNLE/shipping/internals/adapters/internal_adapter/handlers"
)

func PublicRoutesAdaptor(public *echo.Group, handler *handlers.HTTPHandler) *echo.Group {
	public.POST("/register", handler.Register)
	public.POST("/login", handler.Login)
	public.POST("/delivery", handler.DeliveryProduct)
	public.POST("/resetpassword", handler.ResetPassword)
	return public
}
