package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (handler *HTTPHandler) NewAddress(context echo.Context) error {
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
	// To Do
	for i, address := range location.Address {
		externalAddress, err := handler.externalServicesAdapter.TerminalAddressAdaptor(address)
		if err != nil {
			return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
		}
		if externalAddress["data"] != nil {
			result := externalAddress["data"].(map[string]interface{})
			location.Address[i].ExternalAddressID = result["address_id"].(string)
		}
	}
	// Call externalServiceAdapter here
	location.UserID = user.ID
	err = handler.internalServicesAdapter.NewLocationAdaptor(*location);
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusConflict, ADDRESS_ALREADY_EXIST, context)
	}
	// Process valid location data
	return handler.ComputeResponseMessage(http.StatusOK, ADDRESSES_SUBMITTED_SUCCESSFULLY, context)
}

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
