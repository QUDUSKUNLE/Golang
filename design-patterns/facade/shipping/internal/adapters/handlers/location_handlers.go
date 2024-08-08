package handlers

import (
	"net/http"
		"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

func (handler *HTTPHandler) NewAddress(context echo.Context) error {
	location := new(domain.LocationDTO)
	if err := handler.ValidateStruct(context, location); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(),
			context)
	}

	err := handler.servicesAdapter.NewLocationAdaptor(*location);
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
	location, err := handler.servicesAdapter.GetLocationAdaptor(addressID, user.ID);
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusConflict, err.Error(), context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, location, context)
}
