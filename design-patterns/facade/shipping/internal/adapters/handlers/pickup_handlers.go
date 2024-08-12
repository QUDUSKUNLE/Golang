package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

// @Summary Update a pickup
// @Description Update a pickup
// @Tags Carrier Pickup
// @Accept json
// @Produce json
// @Param Body body domain.PickUpDto true "Update a pickup"
// @Param Authorization header string true "Bearer token"
// @Failure 409 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /pickups [put]
func (handler *HTTPHandler) UpdatePickUp(context echo.Context) error {
	pickUpDto := new(domain.PickUpDto)
	if err := handler.ValidateStruct(context, pickUpDto); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err, context)
	}

	// Validate carrier
	carrier, err := handler.ParseUserID(context)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, err.Error(), context)
	}

	if carrier.UserType != string(domain.CARRIER) {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}

	// Initiate a new pick up
	err = handler.internalServicesAdapter.UpDatePickUpAdaptor(*pickUpDto);
	if err != nil {
		if err.Error() == string(RECORD_NOT_FOUND) {
			return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
		}
		return handler.ComputeErrorResponse(http.StatusNotAcceptable, err.Error(), context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, UPDATE_PARCEL_SUCCESSFULLY, context)
}
