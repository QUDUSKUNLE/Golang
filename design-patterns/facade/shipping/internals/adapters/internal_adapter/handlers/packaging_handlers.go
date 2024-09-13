package handlers

import (
	"net/http"
	"sync"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/labstack/echo/v4"
)

// @Summary Submit packagings
// @Description create packagings
// @Tags Packaging
// @Accept json
// @Produce json
// @Param body body domain.TerminalPackagingDto true "Create packagings"
// @Param authorization header string true "Bearer token"
// @Failure 409 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /packagings [post]
func (handler *HTTPHandler) PostPackaging(context echo.Context) error {
	terminalPackaging := new(domain.TerminalPackagingDto)
	if err := ValidateStruct(context, terminalPackaging); err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err, context)
	}
	// Parse user
	user, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return err
	}
	var packagingSync sync.WaitGroup
	packaging := new(domain.PackagingDto)
	packagingChannel := make(chan domain.PackagingDto)
	var externalPackaging map[string]interface{}

	// InternalTerminalAdaptor function
	internalTerminalAdaptor := func (packGroup *sync.WaitGroup, packs domain.PackagingDto) {
		defer packGroup.Done()
		err = handler.internalServicesAdapter.NewPackagingAdaptor(packs)
		if err != nil {
			panic(err)
		}
	}
	// ExternalTerminalAdaptor function
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
	// Close channel function
	closePackagingChannel := func(packGroup *sync.WaitGroup, packagingChannel chan <- domain.PackagingDto) {
		packGroup.Wait()
		close(packagingChannel)
	}
	// Goroutine externalTerminalAdaptor
	for _, pack := range terminalPackaging.Packagings {
		packagingSync.Add(1)
		go externalTerminalAdaptor(&packagingSync, pack, packagingChannel)
	}

	// Goroutine close channel
	go closePackagingChannel(&packagingSync, packagingChannel)

	// Build packaging for internalAdaptor
	packaging.UserID = user.ID
	for c := range packagingChannel {
		packaging.PackagingID = append(packaging.PackagingID, (c.PackagingID)...)
	}
	packagingSync.Add(1)
	go internalTerminalAdaptor(&packagingSync, *packaging)
	packagingSync.Wait()

	// Return response
	return ComputeResponseMessage(http.StatusCreated, PACKAGES_SUBMITTED_SUCCESSFULLY, context)
}
