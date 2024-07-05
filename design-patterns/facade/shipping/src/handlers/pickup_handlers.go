package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/src"
	"github.com/QUDUSKUNLE/shipping/src/dto"
	"github.com/labstack/echo/v4"
)



func PickupProduct(context echo.Context) error {
	pickUpDto := new(dto.PickUpDTO)
	if err := context.Bind(pickUpDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validate shippingDto
	if err := context.Validate(pickUpDto); err != nil {
		return err
	}
	// Initiate a new pick up
	newPickUp := shipping.NewPickUp(pickUpDto.AccountID, pickUpDto.ProductType)

	// Convert ProductType to string
	productType := pickUpDto.ProductType.PrintProduct()

	// Pick up the product
	if err := newPickUp.NewSchedulePickUp(pickUpDto.AccountID, pickUpDto.PickUpAddress, pickUpDto.DeliveryAddress, productType); err != nil {
		return context.JSON(http.StatusNotAcceptable, map[string]string{"message": err.Error(), "success": "false" })
	}
	return context.JSON(http.StatusOK, map[string]string{
		"message": "Product is picked up for shipping.",
		"success": "true",
	})
}
