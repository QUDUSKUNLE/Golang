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
		return context.JSON(http.StatusBadRequest, echo.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}

	err := handler.servicesAdapter.NewLocationAdaptor(*location);
	if err != nil {
		return context.JSON(http.StatusConflict, echo.Map{
			"Message": "Address already exist registered.",
			"Success": false,
		})
	}
	// Process valid location data
	return context.JSON(http.StatusOK, echo.Map{
		"Message": "User addresses submitted successfully.",
		"Success": true,
	})
}

func (handler *HTTPHandler) GetAddress(context echo.Context) error {
	user, err := handler.ParseUserID(context)
	if err != nil {
		return handler.Unauthorized(err.Error(), context)
	}

	if user.UserType != string(domain.USER) {
		return handler.Unauthorized("Unauthorized to perform this operation.", context)
	}

	ID := context.Param("addressID")
	addressID, err := uuid.Parse(ID)
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"Success": false, "Message": err.Error()})
	}
	location, err := handler.servicesAdapter.GetLocationAdaptor(addressID, user.ID);
	if err != nil {
		return context.JSON(http.StatusConflict, echo.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}
	return context.JSON(http.StatusOK, echo.Map{
		"Location": location,
		"Success": true,
	})
}
