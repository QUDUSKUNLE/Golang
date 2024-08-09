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
		return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(),
			context)
	}

	accountID, err := uuid.Parse(deliveryDto.AccountID)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(),
			context)
	}
	// Initiate a new delivery
	if err := handler.servicesAdapter.NewDeliveryAdaptor(accountID, deliveryDto.ProductType); err != nil {
		return handler.ComputeErrorResponse(http.StatusNotAcceptable, err.Error(),
			context)
	}
	// Convert ProductType to string
	productType := deliveryDto.ProductType.PrintProduct()

	// Deliver a product
	if err := handler.servicesAdapter.NewDelivery(accountID, deliveryDto.PickUpAddress, deliveryDto.DeliveryAddress, productType); err != nil {
		return handler.ComputeErrorResponse(http.StatusNotAcceptable, err.Error(),
			context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, PRODUCT_IS_DELIVERED, context)
}
