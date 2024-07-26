package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

func (handler *HTTPHandler) CarrierPickUps(context echo.Context) error {
	user, err := handler.ParseUserID(context)
	if err != nil {
		return context.JSON(http.StatusUnauthorized, echo.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}

	if user.UserType != string(domain.CARRIER) {
		return context.JSON(http.StatusUnauthorized, echo.Map{
			"Message": "Unauthorized to perform this operation.",
			"Success": false,
		})
	}

	pickUps, err := handler.ServicesAdapter.CarrierPickUpsAdaptor(user.ID);
	if err != nil {
		return context.JSON(http.StatusNotImplemented, echo.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}
	return context.JSON(http.StatusOK, echo.Map{
		"Shippings": pickUps,
		"Success": true,
	})
}
