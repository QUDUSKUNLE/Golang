package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/labstack/echo/v4"
)

// @Summary Update a pickup
// @Description update a pickup
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
	if err := ValidateStruct(context, pickUpDto); err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err, context)
	}

	// Validate carrier
	_, err := PrivateMiddlewareContext(context, string(domain.CARRIER))
	if err != nil {
		return err
	}
	// Initiate a new pick up
	err = handler.internalServicesAdapter.UpDatePickUpAdaptor(*pickUpDto);
	if err != nil {
		if err.Error() == string(RECORD_NOT_FOUND) {
			return ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
		}
		return ComputeErrorResponse(http.StatusNotAcceptable, err.Error(), context)
	}
	return ComputeResponseMessage(http.StatusOK, UPDATE_PARCEL_SUCCESSFULLY, context)
}
