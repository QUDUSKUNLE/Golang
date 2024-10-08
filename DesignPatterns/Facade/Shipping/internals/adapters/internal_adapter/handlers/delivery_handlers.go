package handlers

import (
	"github.com/QUDUSKUNLE/shipping/internals/adapters/internal_adapter/dto"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *HTTPHandler) DeliveryProduct(context echo.Context) error {
	deliveryDto := new(dto.DeliveryDto)
	if err := ValidateStruct(context, deliveryDto); err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err,
			context)
	}

	accountID, err := uuid.Parse(deliveryDto.AccountID)
	if err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err.Error(),
			context)
	}
	// Initiate a new delivery
	if err := handler.internalServicesAdapter.NewDeliveryAdaptor(accountID, deliveryDto.ProductType); err != nil {
		return ComputeErrorResponse(http.StatusNotAcceptable, err.Error(),
			context)
	}
	// Convert ProductType to string
	productType := deliveryDto.ProductType.PrintProduct()

	// Deliver a product
	if err := handler.internalServicesAdapter.NewDelivery(accountID, deliveryDto.PickUpAddress, deliveryDto.DeliveryAddress, productType); err != nil {
		return ComputeErrorResponse(http.StatusNotAcceptable, err.Error(),
			context)
	}
	return ComputeResponseMessage(http.StatusOK, PRODUCT_IS_DELIVERED, context)
}
