package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

// @Summary Get carrier pickups
// @Description Get carrier pickups
// @Tags Carrier Pickup
// @Accept json
// @Produce json
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Param Authorization header string true "Bearer token"
// @Router /pickups [get]
func (handler *HTTPHandler) CarrierPickUps(context echo.Context) error {
	user, err := handler.ParseUserID(context)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, err, context)
	}
	if user.UserType != string(domain.CARRIER) {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}

	pickUps, err := handler.internalServicesAdapter.CarrierPickUpsAdaptor(user.ID);
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusNotImplemented, err.Error(), context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, pickUps, context)
}
