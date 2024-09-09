package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/QUDUSKUNLE/shipping/internals/adapters/internal_adapter/handlers"
)

func PublicRoutesAdaptor(public *echo.Group, handler *handlers.HTTPHandler) *echo.Group {
	public.POST("/delivery", handler.DeliveryProduct)
	public.POST("/reject", handler.RejectProduct)
	public.POST("/register", handler.Register)
	public.POST("/login", handler.Login)
	public.POST("/starvation", handler.Starvation)
	public.POST("/example", handler.Example)
	public.POST("/reset_password", handler.ResetPassword)
	public.POST("/live_lock", handler.LiveLock)
	public.POST("/dead_lock", handler.DeadLock)
	return public
}
