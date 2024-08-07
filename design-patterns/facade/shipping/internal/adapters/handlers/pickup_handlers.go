package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

func (handler *HTTPHandler) UpdatePickUp(context echo.Context) error {
	pickUpDto := new(domain.PickUp)
	if err := handler.ValidateStruct(context, pickUpDto); err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}

	// Validate carrier
	carrier, err := handler.ParseUserID(context)
	if err != nil {
		return context.JSON(http.StatusUnauthorized, echo.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}

	if carrier.UserType != string(domain.CARRIER) {
		return context.JSON(http.StatusUnauthorized, echo.Map{
			"Message": "Unauthorized to perform this operation.",
			"Success": false,
		})
	}

	// Initiate a new pick up
	err = handler.servicesAdapter.UpDatePickUpAdaptor(*pickUpDto);
	if err != nil {
		if err.Error() == "record not found" {
			return context.JSON(http.StatusUnauthorized, echo.Map{
				"Message": "User`s unauthorized to perform this operation.",
				"Success": false })
		}
		return context.JSON(http.StatusNotAcceptable, echo.Map{
			"Message": err.Error(),
			"Success": "false",
		})
	}
	return context.JSON(http.StatusOK, echo.Map{
		"Message": "Update parcel successfully.",
		"Success": true,
	})
}
