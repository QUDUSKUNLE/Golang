package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internal/core/services"
	"github.com/QUDUSKUNLE/shipping/internal/core/model"
	"github.com/labstack/echo/v4"
)

func NewShipping(context echo.Context) error {
	shippingDto := new(model.ShippingDTO)
	if err := context.Bind(shippingDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validate shippingDto
	if err := context.Validate(shippingDto); err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"Success": false, "Message": err.Error()})
	}

	// Initiate a new shipping
	err := services.NewShippingAdaptor(context, shippingDto);
	if err != nil {
		return context.JSON(http.StatusNotAcceptable, echo.Map{"Message": err.Error(), "Success": false })
	}
	return context.JSON(http.StatusOK, echo.Map{
		"Message": "Product is scheduled for shipping.",
		"Success": true,
	})
}

func GetShippings(context echo.Context) error {
	shippings, err := services.GetShippingsAdaptor(context);
	if err != nil {
		return context.JSON(http.StatusNotImplemented, echo.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}
	return context.JSON(http.StatusOK, echo.Map{
		"Shippings": shippings,
		"Success": true,
	})
}

func RejectProduct(context echo.Context) error {
	return context.JSON(http.StatusOK, echo.Map{
		"Message": "Product is delivered.",
		"Success": true,
	})
}
