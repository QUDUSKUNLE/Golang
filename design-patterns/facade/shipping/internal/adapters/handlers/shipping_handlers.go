package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

func (handler *HTTPHandler) NewShipping(context echo.Context) error {
	shippingDto := new(domain.ShippingDTO)
	if err := handler.ValidateStruct(context, shippingDto); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(),
			context)
	}

	// Validate carrier
	user, err := handler.ParseUserID(context)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, err.Error(),
		context)
	}

	if user.UserType != string(domain.USER) {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION,
		context)
	}

	shippingDto.UserID = user.ID
	// Initiate a new shipping
	err = handler.servicesAdapter.NewShippingAdaptor(shippingDto);
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusNotAcceptable, err.Error(),
		context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, PRODUCT_SCHEDULED_FOR_SHIPPING, context)
}

func (handler *HTTPHandler) GetShippings(context echo.Context) error {
	user, err := handler.ParseUserID(context)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, err.Error(), context)
	}

	if user.UserType != string(domain.USER) {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}

	shippings, err := handler.servicesAdapter.GetShippingsAdaptor(user.ID);
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusNotImplemented, err.Error(), context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, shippings, context)
}

func (handler *HTTPHandler) RejectProduct(context echo.Context) error {
	return handler.ComputeResponseMessage(http.StatusOK, PRODUCT_IS_DELIVERED, context)
}
