package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

// @Summary Submit a shipment
// @Description create a shipment
// @Tags Shipment
// @Accept json
// @Produce json
// @Param Body body domain.ShippingDto true "Create a shipment"
// @Param Authorization header string true "Bearer token"
// @Failure 409 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /shipments [post]
func (handler *HTTPHandler) PostShipping(context echo.Context) error {
	shippingDto := new(domain.ShippingDto)
	if err := handler.ValidateStruct(context, shippingDto); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err,
			context)
	}
	// Validate carrier
	user, err := handler.PrivateMiddlewareContext(context, string(domain.USER))
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
	err = handler.internalServicesAdapter.NewShippingAdaptor(shippingDto);
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusNotAcceptable, err.Error(),
		context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, PRODUCT_SCHEDULED_FOR_SHIPPING, context)
}

// @Summary Get shipments
// @Description get shipments
// @Tags Shipment
// @Accept json
// @Produce json
// @Failure 401 {object} domain.Response
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} domain.Response
// @Router /shipments [get]
func (handler *HTTPHandler) GetShippings(context echo.Context) error {
	user, err := handler.ParseUserID(context)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, err.Error(), context)
	}

	if user.UserType != string(domain.USER) {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}

	shippings, err := handler.internalServicesAdapter.GetShippingsAdaptor(user.ID);
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusNotImplemented, err.Error(), context)
	}
	return handler.ComputeResponseMessage(http.StatusCreated, shippings, context)
}

func (handler *HTTPHandler) RejectProduct(context echo.Context) error {
	return handler.ComputeResponseMessage(http.StatusOK, PRODUCT_IS_DELIVERED, context)
}
