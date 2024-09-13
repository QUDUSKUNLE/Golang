package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/google/uuid"
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
	terminalShipment := new(domain.TerminalShipmentDto)
	if err := ValidateStruct(context, terminalShipment); err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err,
			context)
	}
	// Validate carrier
	user, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}

	shippingDto := new(domain.ShippingDto)
	// Make call to external adapter to log a shipment
	for _, terminal_shipment := range terminalShipment.Shipments {
		terminal_shipment.ShipmentType = domain.False
		PickUpAddressID, _ := uuid.Parse(terminal_shipment.PickUpAddressID)
		DeliveryAddressID, _ := uuid.Parse(terminal_shipment.DeliveryAddressID)
		terminalPickUpAddress, err := handler.internalServicesAdapter.GetMultipleLocationAdaptor([]uuid.UUID{PickUpAddressID, DeliveryAddressID}, user.ID)
		if err != nil {
			return ComputeErrorResponse(http.StatusNotAcceptable, err.Error(),
		context)
		}
		for _, address := range terminalPickUpAddress {
			if (address.ID == PickUpAddressID) {
				terminal_shipment.PickUpAddressID = address.TerminalAddressID
			}
			if (address.ID == DeliveryAddressID) {
				terminal_shipment.DeliveryAddressID = address.TerminalAddressID
			}
		}
		externalShipment, _ := handler.externalServicesAdapter.TerminalCreateShipmentAdaptor(terminal_shipment)
		if externalShipment["data"] == nil {
			return ComputeErrorResponse(http.StatusBadRequest, externalShipment["message"], context)
		}
		result := externalShipment["data"].(map[string]interface{})
		terminalShipmentID := result["shipment_id"].(string)
		shippingDto.Shipments = append(shippingDto.Shipments, domain.SingleShippingDto{
			PickUpAddressID: PickUpAddressID,
			DeliveryAddressID: DeliveryAddressID,
			UserID: user.ID,
			CarrierID: terminal_shipment.CarrierID,
			Description: terminal_shipment.Description,
			ProductType: terminal_shipment.ProductType,
			TerminalShipmentID: terminalShipmentID,
		})
	}

	err = handler.internalServicesAdapter.NewShippingAdaptor(shippingDto);
	if err != nil {
		return ComputeErrorResponse(http.StatusNotAcceptable, err.Error(),
		context)
	}
	return ComputeResponseMessage(http.StatusOK, PRODUCT_SCHEDULED_FOR_SHIPPING, context)
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
	user, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}

	shippings, err := handler.internalServicesAdapter.GetShippingsAdaptor(user.ID);
	if err != nil {
		return ComputeErrorResponse(http.StatusNotImplemented, err.Error(), context)
	}
	return ComputeResponseMessage(http.StatusOK, shippings, context)
}
