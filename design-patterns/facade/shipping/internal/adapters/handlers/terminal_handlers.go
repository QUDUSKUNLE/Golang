package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

func (handler *HTTPHandler) Rates(context echo.Context) error {
	response, err := handler.externalServicesAdapter.TerminalGetRatesAdaptor()
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusOK, err, context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, response, context)
}

func (handler *HTTPHandler) Packaging(context echo.Context) error {
	packaging := new(domain.PackagingDTO)
	if err := handler.ValidateStruct(context, packaging); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	response, err := handler.externalServicesAdapter.TerminalPackagingAdaptor(*packaging)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusOK, err, context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, response, context)
}
