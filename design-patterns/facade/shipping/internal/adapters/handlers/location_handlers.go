package handlers

import (
	"net/http"

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
		if err.Error() == "location`s already exist" {
			return context.JSON(http.StatusConflict, echo.Map{
				"Message": "User already registered",
				"Success": false })
		}
		if err.Error() == `incorrect passwords` {
			return context.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Success": false,
			})
		}
		return context.JSON(http.StatusConflict, echo.Map{
			"Message": "User`s already registered.",
			"Success": false,
		})
	}
	// Process valid location data
	return context.JSON(http.StatusOK, echo.Map{
		"Message": "User registered successfully.",
		"Success": true,
	})
}
