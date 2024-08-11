package handlers

import (
	"fmt"
	"net/http"

	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

func (handler *HTTPHandler) PostParcel(context echo.Context) error {
	terminalParcel := new(domain.TerminalParcelDto)
	if err := handler.ValidateStruct(context, terminalParcel); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err, context)
	}
	// Validate user
	user, err := handler.ParseUserID(context)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, err.Error(), context)
	}

	if user.UserType != string(domain.USER) {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}
	parcel := new(domain.ParcelDto)
	// Make call to external adapter to register parcel
	for _, terminal_parcel := range terminalParcel.Parcels {
		externalParcel, err := handler.externalServicesAdapter.TerminalCreateParcelAdaptor(terminal_parcel)
		if err != nil {
			return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
		}
		if externalParcel["data"] != nil {
			result := externalParcel["data"].(map[string]interface{})
			parcel_id := result["parcel_id"].(string)
			parcel.ParcelID = append(parcel.ParcelID, parcel_id)
		} else {
				fmt.Print(externalParcel["data"])
				return handler.ComputeErrorResponse(http.StatusBadRequest, externalParcel["message"], context)
		}
	}
	parcel.UserID = user.ID
	if err := handler.internalServicesAdapter.NewParcelAdaptor(*parcel); err != nil {
		return handler.ComputeErrorResponse(http.StatusConflict, "Parcel error", context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, PARCEL_SUBMITTED_SUCCESSFULLY, context)
}
