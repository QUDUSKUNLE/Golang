package handlers

import (
	"net/http"

	"github.com/google/uuid"
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
		return context.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": err.Error() })
	}

	accountID, err := uuid.Parse(shippingDto.AccountID);
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": err.Error() })
	}
	// Initiate new shipping
	newShipping := shipping.NewShippingAdaptor(accountID, shippingDto.ProductType)

	// Convert ProductType to string
	productType := shippingDto.ProductType.PrintProduct()

	// Schedule shipping
	if err := newShipping.NewShipping(accountID, shippingDto.PickUpAddress, shippingDto.DeliveryAddress, productType); err != nil {
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
