package handlers

import (
	// "fmt"
	"net/http"

	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)


// @Summary Submit packagings
// @Description Create packagings
// @Tags Packaging
// @Accept json
// @Produce json
// @Param Body body domain.TerminalPackagingDto true "Create packagings"
// @Param Authorization header string true "Bearer token"
// @Failure 409 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /packagings [post]
func (handler *HTTPHandler) PostPackaging(context echo.Context) error {
	terminalPackaging := new(domain.TerminalPackagingDto)
	if err := handler.ValidateStruct(context, terminalPackaging); err != nil {
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
	packaging := new(domain.PackagingDto)
	// Make call to external adapter to register packaging
	for _, pack := range terminalPackaging.Packagings {
		externalPackaging, err := handler.externalServicesAdapter.TerminalCreatePackagingAdaptor(pack)
		if err != nil {
			return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
		}
		if externalPackaging["data"] != nil {
			result := externalPackaging["data"].(map[string]interface{})
			packaging_id := result["packaging_id"].(string)
			packaging.PackagingID = append(packaging.PackagingID, packaging_id)
		} else {
			return handler.ComputeErrorResponse(http.StatusBadRequest, externalPackaging["message"], context)
		}
	}
	packaging.UserID = user.ID
	err = handler.internalServicesAdapter.NewPackagingAdaptor(*packaging)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusConflict, "Package error", context)
	}
	return handler.ComputeResponseMessage(http.StatusCreated, PACKAGES_SUBMITTED_SUCCESSFULLY, context)
}
