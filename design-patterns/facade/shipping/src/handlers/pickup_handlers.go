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
			"message": err.Error(),
			"success": false,
		})
	}
	// Validate pickUpDto
	if err := context.Validate(pickUpDto); err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": err.Error() })
	}
	// Initiate a new pick up
	err := shipping.UpDatePickUpAdaptor(*pickUpDto);
	if err != nil {
		return context.JSON(http.StatusNotAcceptable, echo.Map{"message": err.Error(), "success": "false" })
	}
	return context.JSON(http.StatusOK, echo.Map{
		"message": "Update parcel successfully.",
		"success": true,
	})
}
