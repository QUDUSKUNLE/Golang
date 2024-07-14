package handlers

import (
	"net/http"
	"fmt"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

func (handler *HTTPHandler) UpdatePickUp(context echo.Context) error {
	pickUpDto := new(domain.PickUp)
	if err := handler.ValidateStruct(context, pickUpDto); err != nil {
		return err
	}

	// Validate carrier
	carrier, err := handler.ParseUserID(context)
	if err != nil {
		return err
	}

	if carrier.UserType != string(domain.RIDER) {
		return fmt.Errorf("unauthorized to perform this operation")
	}

	// Initiate a new pick up
	err = handler.ServicesAdapter.UpDatePickUpAdaptor(*pickUpDto);
	if err != nil {
		if err.Error() == "record not found" {
			return context.JSON(http.StatusUnauthorized, echo.Map{"Message": "User`s unauthorized to perform this operation.", "Success": false })
		}
		return context.JSON(http.StatusNotAcceptable, echo.Map{"Message": err.Error(), "Success": "false" })
	}
	return context.JSON(http.StatusOK, echo.Map{
		"Message": "Update parcel successfully.",
		"Success": true,
	})
}
