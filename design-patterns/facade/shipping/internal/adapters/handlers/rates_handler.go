package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func (handler *HTTPHandler) Rates(context echo.Context) error {
	response, err := handler.externalServicesAdapter.TerminalGetRatesAdaptor()
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusOK, err, context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, response, context)
}
