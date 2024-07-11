package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/src"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/labstack/echo/v4"
)

func UpdatePickUp(context echo.Context) error {
	pickUpDto := new(model.PickUp)
	if err := context.Bind(pickUpDto); err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}
	// Validate pickUpDto
	if err := context.Validate(pickUpDto); err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"Success": false, "Message": err.Error() })
	}
	// Initiate a new pick up
	err := shipping.UpDatePickUpAdaptor(context, *pickUpDto);
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
