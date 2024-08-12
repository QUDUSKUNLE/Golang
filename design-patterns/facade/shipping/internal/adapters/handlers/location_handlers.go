package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// @Summary Submit addresses
// @Description create addresses
// @Tags Address
// @Accept json
// @Produce json
// @Param Body body domain.LocationDto true "Create addresses"
// @Param Authorization header string true "Bearer token"
// @Failure 409 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /addresses [post]
func (handler *HTTPHandler) PostAddress(context echo.Context) error {
	location := new(domain.LocationDto)
	if err := handler.ValidateStruct(context, location); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err,
			context)
	}

	user, err := handler.ParseUserID(context)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, err.Error(), context)
	}

	if user.UserType != string(domain.USER) {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}
	// Make call to external adapter to register address
	for i, address := range location.Address {
		externalAddress, _ := handler.externalServicesAdapter.TerminalCreateAddressAdaptor(address)
		if externalAddress["data"] != nil {
			result := externalAddress["data"].(map[string]interface{})
			address_id := result["address_id"].(string)
			location.Address[i].TerminalAddressID = address_id
		} else {
			return handler.ComputeErrorResponse(http.StatusBadRequest, externalAddress["message"], context)
		}
	}
	// Make call to internal adapter to save register
	location.UserID = user.ID
	err = handler.internalServicesAdapter.NewLocationAdaptor(*location);
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusConflict, ADDRESS_ALREADY_EXIST, context)
	}
	// Process valid location data
	return handler.ComputeResponseMessage(http.StatusCreated, ADDRESSES_SUBMITTED_SUCCESSFULLY, context)
}

// @Summary Get a address
// @Description get a address
// @Tags Address
// @Accept json
// @Produce json
// @Param addressID path string true "Address ID"
// @Param Authorization header string true "Bearer token"
// @Failure 400 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /addresses/:addressID [get]
func (handler *HTTPHandler) GetAddress(context echo.Context) error {
	user, err := handler.ParseUserID(context)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, err.Error(), context)
	}

	if user.UserType != string(domain.USER) {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}

	ID := context.Param("addressID")
	addressID, err := uuid.Parse(ID)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	location, err := handler.internalServicesAdapter.GetLocationAdaptor(addressID, user.ID);
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusConflict, err.Error(), context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, location, context)
}

// @Summary Get addresses
// @Description get addresses
// @Tags Address
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Failure 400 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /addresses [get]
func (handler *HTTPHandler) GetAddresses(context echo.Context) error {
	user, err := handler.ParseUserID(context)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, err.Error(), context)
	}

	if user.UserType != string(domain.USER) {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}

	locations, err := handler.internalServicesAdapter.GetLocationsAdaptor(user.ID);
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusConflict, err.Error(), context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, locations, context)
}
