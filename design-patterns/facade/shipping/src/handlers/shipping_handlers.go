package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/src"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/labstack/echo/v4"
)

func NewShipping(context echo.Context) error {
	shippingDto := new(model.ShippingDTO)
	if err := context.Bind(shippingDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validate shippingDto
	if err := context.Validate(shippingDto); err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": err.Error()})
	}

	// Initiate a new shipping
	err := shipping.NewShippingAdaptor(context, shippingDto);
	if err != nil {
		return context.JSON(http.StatusNotAcceptable, echo.Map{"message": err.Error(), "success": false })
	}
	return context.JSON(http.StatusOK, echo.Map{
		"message": "Product is scheduled for shipping.",
		"success": true,
	})
}

func GetShippings(context echo.Context) error {
	shippings, err := shipping.GetShippingsAdaptor(context);
	if err != nil {
		return context.JSON(http.StatusNotImplemented, echo.Map{
			"message": err.Error(),
			"success": false,
		})
	}
	return context.JSON(http.StatusOK, echo.Map{
		"shippings": shippings,
		"success": true,
	})
}

func RejectProduct(context echo.Context) error {
	return context.JSON(http.StatusOK, echo.Map{
		"message": "Product is delivered.",
		"success": true,
	})
}
