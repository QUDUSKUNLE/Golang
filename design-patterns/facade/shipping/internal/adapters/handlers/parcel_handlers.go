package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

func (handler *HTTPHandler) PostParcel(context echo.Context) error {
	parcel := new(domain.TerminalParcelDto)
	if err := handler.ValidateStruct(context, parcel); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err, context)
	}
	response, err := handler.externalServicesAdapter.TerminalCreateParcelAdaptor(*parcel)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusOK, err, context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, response, context)
}
