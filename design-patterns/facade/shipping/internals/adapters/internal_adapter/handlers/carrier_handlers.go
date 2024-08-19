package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/labstack/echo/v4"
)

// @Summary Get carrier pickups
// @Description get carrier pickups
// @Tags Carrier Pickup
// @Accept json
// @Produce json
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Param Authorization header string true "Bearer token"
// @Router /pickups [get]
func (handler *HTTPHandler) CarrierPickUps(context echo.Context) error {
	user, err := PrivateMiddlewareContext(context, string(domain.CARRIER))
	if err != nil {
		return ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}

	pickUps, err := handler.internalServicesAdapter.CarrierPickUpsAdaptor(user.ID);
	if err != nil {
		return ComputeErrorResponse(http.StatusNotImplemented, err.Error(), context)
	}
	return ComputeResponseMessage(http.StatusOK, pickUps, context)
}
