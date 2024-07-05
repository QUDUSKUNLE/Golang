package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/src"
	"github.com/QUDUSKUNLE/shipping/src/dto"
	"github.com/labstack/echo/v4"
)

func ScheduleProduct(context echo.Context) error {
	shippingDto := new(dto.ShippingDTO)
	if err := context.Bind(shippingDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validate shippingDto
	if err := context.Validate(shippingDto); err != nil {
		return err
	}
	// Initiate new shipping
	newShipping := shipping.NewShipping(shippingDto.AccountID, shippingDto.ProductType)

	// Convert ProductType to string
	productType := shippingDto.ProductType.PrintProduct()

	// Schedule shipping
	if err := newShipping.NewScheduleShipping(shippingDto.AccountID, shippingDto.PickUpAddress, shippingDto.DeliveryAddress, productType); err != nil {
		return context.JSON(http.StatusNotAcceptable, map[string]string{"message": err.Error(), "success": "false" })
	}
	return context.JSON(http.StatusOK, map[string]string{
		"message": "Product is scheduled for shipping.",
		"success": "true",
	})
}

func RejectProduct(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]string{
		"message": "Product is delivered.",
		"success": "true",
	})
}
