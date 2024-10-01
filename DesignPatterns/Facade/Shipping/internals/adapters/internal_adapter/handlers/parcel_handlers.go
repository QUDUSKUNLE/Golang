package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/labstack/echo/v4"
)

// @Summary Submit parcels
// @Description create parcels
// @Tags Parcel
// @Accept json
// @Produce json
// @Param body body domain.TerminalParcelDto true "Create parcels"
// @Param authorization header string true "Bearer token"
// @Failure 409 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /parcels [post]
func (handler *HTTPHandler) PostParcel(context echo.Context) error {
	terminalParcel := new(domain.TerminalParcelDto)
	if err := ValidateStruct(context, terminalParcel); err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err, context)
	}
	// Validate user
	user, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return err
	}

	parcel := new(domain.ParcelDto)
	// Make call to external adapter to register parcel
	for _, terminal_parcel := range terminalParcel.Parcels {
		externalParcel, err := handler.externalServicesAdapter.TerminalCreateParcelAdaptor(terminal_parcel)
		if err != nil {
			return ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
		}
		if externalParcel["data"] != nil {
			result := externalParcel["data"].(map[string]interface{})
			parcel_id := result["parcel_id"].(string)
			parcel.ParcelID = append(parcel.ParcelID, parcel_id)
		} else {
				return ComputeErrorResponse(http.StatusBadRequest, externalParcel["message"], context)
		}
	}
	parcel.UserID = user.ID
	if err := handler.internalServicesAdapter.NewParcelAdaptor(*parcel); err != nil {
		return ComputeErrorResponse(http.StatusConflict, "Parcel error", context)
	}
	return ComputeResponseMessage(http.StatusCreated, PARCEL_SUBMITTED_SUCCESSFULLY, context)
}

// @Summary Get parcels
// @Description get parcels
// @Tags Parcel
// @Accept json
// @Produce json
// @Param authorization header string true "Bearer token"
// @Failure 400 {object} domain.Response
// @Success 200 {object} domain.Response
// @Router /parcels [get]
func (handler *HTTPHandler) GetParcels(context echo.Context) error {
	user, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return err
	}
	parcels, err := handler.internalServicesAdapter.GetParcelsAdaptor(user.ID); 		if err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, "Parcel error", context)
	}
	return ComputeResponseMessage(http.StatusOK, parcels, context)
}


// @Summary Get a parcel
// @Description get a parcel
// @Tags Parcel
// @Accept json
// @Produce json
// @Param authorization header string true "Bearer token"
// @Param parcel_id path string true "Parcel ID"
// @Failure 400 {object} domain.Response
// @Success 200 {object} domain.Response
// @Router /parcels/{parcel_id} [get]
func (handler *HTTPHandler) GetParcel(context echo.Context) error {
	user, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return err
	}
	parcels, err := handler.internalServicesAdapter.GetParcelsAdaptor(user.ID); 		if err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, "Parcel error", context)
	}
	return ComputeResponseMessage(http.StatusOK, parcels, context)
}

// @Summary Update a parcel
// @Description update a parcel
// @Tags Parcel
// @Accept json
// @Produce json
// @Param authorization header string true "Bearer token"
// @Param parcel_id path string true "Parcel ID"
// @Param body body domain.LocationDto true "Update a parcel"
// @Failure 400 {object} domain.Response
// @Success 200 {object} domain.Response
// @Router /parcels/{parcel_id} [put]
func (handler *HTTPHandler) PutParcel(context echo.Context) error {
	_, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}
	description := context.QueryParam("description")
	
	return ComputeResponseMessage(http.StatusOK, description, context)
}

// @Summary Delete a parcel
// @Description delete a parcel
// @Tags Parcel
// @Accept json
// @Produce json
// @Param authorization header string true "Bearer token"
// @Param parcel_id path string true "Parcel ID"
// @Failure 400 {object} domain.Response
// @Success 204 {object} domain.Response
// @Router /parcels/{parcel_id} [delete]
func (handler *HTTPHandler) DeleteParcel(context echo.Context) error {
	_, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}
	description := context.QueryParam("description")
	
	return ComputeResponseMessage(http.StatusOK, description, context)
}
