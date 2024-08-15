package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/QUDUSKUNLE/shipping/internal/adapters/handlers"
)

func PublicRoutesAdaptor(public *echo.Group, handler *handlers.HTTPHandler) *echo.Group {
	public.POST("/delivery", handler.DeliveryProduct)
	public.POST("/reject", handler.RejectProduct)
	public.POST("/register", handler.Register)
	public.POST("/login", handler.Login)
	public.POST("/reset_password", handler.ResetPassword)
	public.POST("/livelock", handler.LiveLock)
	public.POST("/deadlock", handler.DeadLock)
	public.POST("/starvation", handler.Starvation)
	public.POST("/example", handler.Example)
	return public
}
