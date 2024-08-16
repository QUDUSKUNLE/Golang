package handlers

import (
	"net/http"
	"sync"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

// @Summary Submit packagings
// @Description create packagings
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
	var packagingWaitGroup sync.WaitGroup
	packaging := new(domain.PackagingDto)
	packagingChan := make(chan domain.PackagingDto)
	var externalPackaging map[string]interface{}
	// Make call to external adapter to register packaging
	externalTerminalAdaptor := func(packGroup *sync.WaitGroup, pack domain.SingleTerminalPackagingDto, packagingChannel chan <- domain.PackagingDto) {
		defer packGroup.Done()
		externalPackaging, err = handler.externalServicesAdapter.TerminalCreatePackagingAdaptor(pack)
		if err != nil {
			panic(err)
		}
		if externalPackaging["data"] != nil {
			result := externalPackaging["data"].(map[string]interface{})
			packaging_id := result["packaging_id"].(string)
			packs := &domain.PackagingDto{
				PackagingID: []string{packaging_id},
			}
			packagingChannel <- *packs
		}
	}
	internalTerminalAdaptor := func (packGroup *sync.WaitGroup, packs domain.PackagingDto) {
		defer packGroup.Done()
		err = handler.internalServicesAdapter.NewPackagingAdaptor(packs)
		if err != nil {
			panic(err)
		}
	}
	for _, pack := range terminalPackaging.Packagings {
		packagingWaitGroup.Add(1)
		go externalTerminalAdaptor(&packagingWaitGroup, pack, packagingChan)
	}
	closeChannel := func(packGroup *sync.WaitGroup, packagingChannel chan <- domain.PackagingDto) {
		packGroup.Wait()
		close(packagingChannel)
	}

	// Close channel
	go closeChannel(&packagingWaitGroup, packagingChan)

	// Build packaging for internalAdaptor
	packaging.UserID = user.ID
	for c := range packagingChan {
		packaging.PackagingID = append(packaging.PackagingID, (c.PackagingID)...)
	}
	packagingWaitGroup.Add(1)
	go internalTerminalAdaptor(&packagingWaitGroup, *packaging)
	packagingWaitGroup.Wait()
	return handler.ComputeResponseMessage(http.StatusCreated, PACKAGES_SUBMITTED_SUCCESSFULLY, context)
}
