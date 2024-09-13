package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
)

// @Summary Update a pickup
// @Description update a pickup
// @Tags Carrier Pickup
// @Accept json
// @Produce json
// @Param body body domain.PickUpDto true "Update a pickup"
// @Param authorization header string true "Bearer token"
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
		return ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
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

// @Summary Get a pickup
// @Description get a pickup
// @Tags Carrier Pickup
// @Accept json
// @Produce json
// @Param pick_up_id path string true "PickUp ID"
// @Param authorization header string true "Bearer token"
// @Failure 400 {object} domain.Response
// @Success 200 {object} domain.Response
// @Router /pickups/{pick_up_id} [get]
func (handler *HTTPHandler) GetPickUp(context echo.Context) error {
	pickUpParam := context.Param("pick_up_id")
	if pickUpParam == "" {
		return ComputeErrorResponse(http.StatusBadRequest, "Pick_up_id is rquired", context)
	}
	pickUpUUID, err := uuid.Parse(pickUpParam)
	if err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	// Validate carrier
	user, err := PrivateMiddlewareContext(context, string(domain.CARRIER))
	if err != nil {
		return ComputeErrorResponse(http.StatusUnauthorized, err.Error(), context)
	}
	// Initiate a new pick up
	pickUp, err := handler.internalServicesAdapter.GetPickUpAdaptor(pickUpUUID, user.ID);
	if err != nil {
		return ComputeErrorResponse(http.StatusNotAcceptable, err.Error(), context)
	}
	return ComputeResponseMessage(http.StatusOK, pickUp, context)
}
