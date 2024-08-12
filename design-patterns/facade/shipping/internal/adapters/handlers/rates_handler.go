package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

// @Summary Get parcel rates
// @Description Get parcel rates
// @Tags Parcel
// @Accept json
// @Produce json
// @Failure 409 {object} domain.Response
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} domain.Response
// @Router /rates [get]
func (handler *HTTPHandler) Rates(context echo.Context) error {
	response, err := handler.externalServicesAdapter.TerminalGetRatesAdaptor()
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusOK, err, context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, response, context)
}
