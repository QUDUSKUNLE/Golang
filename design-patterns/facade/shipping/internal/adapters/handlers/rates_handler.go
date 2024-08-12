package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

// @Summary Get rates
// @Description get rates
// @Tags Rates
// @Accept json
// @Produce json
// @Failure 400 {object} domain.Response
// @Param Authorization header string true "Bearer Token"
// @Param currency query string true "Currency" Enums(NGN, USD)
// @Param pickup_address_id query string true "Pickup Address ID"
// @Param delivery_address_id query string true "Delivery Address ID"
// @Param parcel_id query string true "Parcel ID"
// @Param cash_on_delivery query string true "Cash On Delivery" Enums(false, true)
// @Success 200 {object} domain.Response
// @Router /rates [get]
func (handler *HTTPHandler) Rates(context echo.Context) error {
	// Validate user
	rates := &domain.TerminalRatesQueryDto{
		Currency: domain.Currency(context.QueryParam("currency")),
		PickUpAddressID: context.QueryParam("pickup_address_id"),
		DeliveryAddressID: context.QueryParam("delivery_address_id"),
		ParcelID: context.QueryParam("parcel_id"),
		CashOnDelivery: domain.CASH_ON_DELIVERY(context.QueryParam("cash_on_delivery")),
	}
	if rates.Currency == "" {
		return handler.ComputeErrorResponse(http.StatusBadRequest, "QueryParam currency is rquired", context)
	}
	if rates.PickUpAddressID == "" {
		return handler.ComputeErrorResponse(http.StatusBadRequest, "QueryParam pickup_address_id is rquired", context)
	}
	if rates.DeliveryAddressID == "" {
		return handler.ComputeErrorResponse(http.StatusBadRequest, "QueryParam delivery_address_id is rquired", context)
	}
	if rates.ParcelID == "" {
		return handler.ComputeErrorResponse(http.StatusBadRequest, "QueryParam parcel_id is rquired", context)
	}
	user, err := handler.ParseUserID(context)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, err.Error(), context)
	}

	if user.UserType != string(domain.USER) {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}
	response, err := handler.externalServicesAdapter.TerminalGetRatesAdaptor(*rates)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusOK, err, context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, response, context)
}
