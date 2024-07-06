package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/src"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func ScheduleShipping(context echo.Context) error {
	shippingDto := new(model.ShippingDTO)
	if err := context.Bind(shippingDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validate shippingDto
	if err := context.Validate(shippingDto); err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": err.Error()})
	}
	// Initiate new shipping
	newShipping := shipping.NewShippingAdaptor()
	// Parse ID
	ID, err := uuid.Parse(newShipping.Utils.ObtainUser(context));
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": err.Error() })
	}

	// Schedule shipping
	if err := newShipping.NewShipping(ID, *shippingDto); err != nil {
		return context.JSON(http.StatusNotAcceptable, echo.Map{"message": err.Error(), "success": false })
	}
	return context.JSON(http.StatusOK, echo.Map{
		"message": "Product is scheduled for shipping.",
		"success": true,
	})
}

func RejectProduct(context echo.Context) error {
	return context.JSON(http.StatusOK, echo.Map{
		"message": "Product is delivered.",
		"success": true,
	})
}
