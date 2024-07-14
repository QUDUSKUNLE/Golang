package handlers

import (
	"github.com/QUDUSKUNLE/shipping/internal/adapters/dto"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *HTTPHandler) DeliveryProduct(context echo.Context) error {
	deliveryDto := new(dto.DeliveryDTO)
	if err := handler.ValidateStruct(context, deliveryDto); err != nil {
		return err
	}

	accountID, err := uuid.Parse(deliveryDto.AccountID)
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"Success": false, "Message": err.Error()})
	}
	// Initiate a new delivery
	newDelivery := handler.ServicesAdapter.NewDeliveryAdaptor(accountID, deliveryDto.ProductType)

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
