package handlers

import (
	"github.com/QUDUSKUNLE/shipping/src"
	"github.com/QUDUSKUNLE/shipping/src/dto"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func DeliveryProduct(context echo.Context) error {
	deliveryDto := new(dto.DeliveryDTO)
	if err := context.Bind(deliveryDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validate shippingDto
	if err := context.Validate(deliveryDto); err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"Success": false, "Message": err.Error()})
	}

	accountID, err := uuid.Parse(deliveryDto.AccountID)
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"Success": false, "Message": err.Error()})
	}
	// Initiate a new delivery
	newDelivery := shipping.NewDeliveryAdaptor(accountID, deliveryDto.ProductType)

	// Convert ProductType to string
	productType := deliveryDto.ProductType.PrintProduct()

	// Deliver a product
	if err := newDelivery.NewDelivery(accountID, deliveryDto.PickUpAddress, deliveryDto.DeliveryAddress, productType); err != nil {
		return context.JSON(http.StatusNotAcceptable, echo.Map{"Message": err.Error(), "Success": false})
	}
	return context.JSON(http.StatusOK, echo.Map{
		"Message": "Product is delivered.",
		"Success": true,
	})
}
