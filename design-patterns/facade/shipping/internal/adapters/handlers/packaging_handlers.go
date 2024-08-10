package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

func (handler *HTTPHandler) PostPackaging(context echo.Context) error {
	packaging := new(domain.TerminalPackagingDto)
	if err := handler.ValidateStruct(context, packaging); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err, context)
	}
	response, err := handler.externalServicesAdapter.TerminalCreatePackagingAdaptor(*packaging)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusOK, err, context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, response, context)
}
