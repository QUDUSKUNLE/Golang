package handlers

import (
	"net/http"
	"sync"

	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// @Summary Submit addresses
// @Description create addresses
// @Tags Address
// @Accept json
// @Produce json
// @Param body body domain.LocationDto true "Create addresses"
// @Param authorization header string true "Bearer token"
// @Failure 409 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /addresses [post]
func (handler *HTTPHandler) PostAddress(context echo.Context) error {
	location := new(domain.LocationDto)
	if err := ValidateStruct(context, location); err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err,
			context)
	}

	// Parse user
	user, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}

	// Make call to internal adapter to save register
	location.UserID = user.ID
	err = handler.internalServicesAdapter.NewLocationAdaptor(*location);
	if err != nil {
		return ComputeErrorResponse(http.StatusConflict, ADDRESS_ALREADY_EXIST, context)
	}

	// Need to run this with goroutine, working on this
	var addressSync sync.WaitGroup
	var externalAddress map[string]interface{}
	var terminalLocation *domain.Location
	terminalLocationChannel := make(chan domain.Location)

	// InternalTerminalAdaptor function
	internalTerminalAdaptor := func (addSync *sync.WaitGroup, location domain.Location) {
		defer addSync.Done()
		err := handler.internalServicesAdapter.TerminalUpdateAddressAdaptor(location)
		if err != nil {
			panic(err)
		}
	}
	// ExternalTerminalAdaptor function
	externalTerminalAdaptor := func (addSync *sync.WaitGroup, address domain.Address, index int, location domain.LocationDto, locationChannel chan <- domain.Location)   {
		defer addSync.Done()
		externalAddress, _ = handler.externalServicesAdapter.TerminalCreateAddressAdaptor(address)
		if externalAddress["data"] != nil {
			result := externalAddress["data"].(map[string]interface{})
			address_id := result["address_id"].(string)
			location.Address[index].TerminalAddressID = address_id
			terminalLocation = &domain.Location{
				TerminalAddressID: address_id,
				Address: location.Address[index],
				Description: location.Address[index].Description,
				UserID: location.UserID,
			}
			// Send terminalLocation to the channel
			locationChannel <- *terminalLocation
		}
	}
	// Close channel function
	closeTerminalChannel := func(terminalGroup *sync.WaitGroup, channel chan <- domain.Location) {
		terminalGroup.Wait()
		close(channel)
	}
	// Goroutine externalTerminalAdaptor
	for index, address := range location.Address {
		addressSync.Add(1)
		go externalTerminalAdaptor(&addressSync, address, index, *location, terminalLocationChannel)
	}
	// Goroutine close channel
	go closeTerminalChannel(&addressSync, terminalLocationChannel)

	// Goroutine internalTerminalAdaptor
	for terminalChannel := range terminalLocationChannel {
		addressSync.Add(1)
		go internalTerminalAdaptor(&addressSync, terminalChannel)
	}
	addressSync.Wait()

	// Process valid location data
	return ComputeResponseMessage(http.StatusCreated, ADDRESSES_SUBMITTED_SUCCESSFULLY, context)
}

// @Summary Get a address
// @Description get a address
// @Tags Address
// @Accept json
// @Produce json
// @Param address_id path string true "Address ID"
// @Param authorization header string true "Bearer token"
// @Failure 400 {object} domain.Response
// @Success 200 {object} domain.Response
// @Router /addresses/{address_id} [get]
func (handler *HTTPHandler) GetAddress(context echo.Context) error {
	user, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return err
	}

	ID := context.Param("address_id")
	addressID, err := uuid.Parse(ID)
	if err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	location, err := handler.internalServicesAdapter.GetLocationAdaptor(addressID, user.ID);
	if err != nil {
		return ComputeErrorResponse(http.StatusConflict, err.Error(), context)
	}
	return ComputeResponseMessage(http.StatusOK, location, context)
}

// @Summary Get addresses
// @Description get addresses
// @Tags Address
// @Accept json
// @Produce json
// @Param authorization header string true "Bearer token"
// @Param description query string false "Description"
// @Failure 400 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /addresses [get]
func (handler *HTTPHandler) GetAddresses(context echo.Context) error {
	user, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}
	description := context.QueryParam("description")
	if description == "" {
		locations, err := handler.internalServicesAdapter.GetLocationsAdaptor(user.ID);
		if err != nil {
			return ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
		}
		return ComputeResponseMessage(http.StatusOK, locations, context)
	} else {
		location, err := handler.internalServicesAdapter.QueryLocationAdaptor(user.ID, description);
			if err != nil {
				return ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
			}
			return ComputeResponseMessage(http.StatusOK, location, context)
	}
}

// @Summary Update an address
// @Description update an address
// @Tags Address
// @Accept json
// @Produce json
// @Param address_id path string true "Address ID"
// @Param body body domain.LocationDto true "Update an address"
// @Param authorization header string true "Bearer token"
// @Failure 400 {object} domain.Response
// @Success 200 {object} domain.Response
// @Router /addresses/{address_id} [put]
func (handler *HTTPHandler) UpdateAddress(context echo.Context) error {
	_, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}
	description := context.QueryParam("description")
	
	return ComputeResponseMessage(http.StatusOK, description, context)
}

// @Summary Delete an address
// @Description delete an address
// @Tags Address
// @Accept json
// @Produce json
// @Param address_id path string true "Address ID"
// @Param authorization header string true "Bearer token"
// @Failure 400 {object} domain.Response
// @Success 204 {object} domain.Response
// @Router /addresses/{address_id} [delete]
func (handler *HTTPHandler) DeleteAddress(context echo.Context) error {
	_, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}
	description := context.QueryParam("description")
	
	return ComputeResponseMessage(http.StatusOK, description, context)
}
