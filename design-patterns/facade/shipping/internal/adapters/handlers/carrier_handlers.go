package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

func (handler *HTTPHandler) CarrierPickUps(context echo.Context) error {
	user, err := handler.ParseUserID(context)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, err.Error(), context)
	}

	if user.UserType != string(domain.CARRIER) {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, "Unauthorized to perform this operation.", context)
	}

	pickUps, err := handler.servicesAdapter.CarrierPickUpsAdaptor(user.ID);
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusNotImplemented, err.Error(), context)
	}

	return context.JSON(http.StatusOK, echo.Map{
		"PickUps": pickUps,
		"Success": true,
	})
}
