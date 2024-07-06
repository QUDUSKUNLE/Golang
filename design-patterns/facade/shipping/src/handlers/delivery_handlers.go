package handlers

import (
	"net/http"
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/src"
	"github.com/QUDUSKUNLE/shipping/src/dto"
	"github.com/labstack/echo/v4"
)


func DeliveryProduct(context echo.Context) error {
	deliveryDto := new(dto.DeliveryDTO)
	if err := context.Bind(deliveryDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validate shippingDto
	if err := context.Validate(deliveryDto); err != nil {
		return err
	}

	accountID, err := uuid.Parse(deliveryDto.AccountID);
	if err != nil {
		return err
	}
	// Initiate a new delivery
	newDelivery := shipping.NewDeliveryAdaptor(accountID, deliveryDto.ProductType)

	// Convert ProductType to string
	productType := deliveryDto.ProductType.PrintProduct()

	// Deliver a product
	if err := newDelivery.NewDelivery(accountID, deliveryDto.PickUpAddress, deliveryDto.DeliveryAddress,productType); err != nil {
		return context.JSON(http.StatusNotAcceptable, map[string]string{"message": err.Error(), "success": "false" })
	}
	return context.JSON(http.StatusOK, map[string]string{
		"message": "Product is delivered.",
		"success": "true",
	})
}
