package handlers

import (
	"fmt"
	"sync"
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var wg sync.WaitGroup

// @Summary Submit addresses
// @Description create addresses
// @Tags Address
// @Accept json
// @Produce json
// @Param Body body domain.LocationDto true "Create addresses"
// @Param Authorization header string true "Bearer token"
// @Failure 409 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /addresses [post]
func (handler *HTTPHandler) PostAddress(context echo.Context) error {
	location := new(domain.LocationDto)
	if err := handler.ValidateStruct(context, location); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err,
			context)
	}

	user, err := handler.ParseUserID(context)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, err.Error(), context)
	}

	if user.UserType != string(domain.USER) {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}

	// Make call to internal adapter to save register
	location.UserID = user.ID
	err = handler.internalServicesAdapter.NewLocationAdaptor(*location);
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusConflict, ADDRESS_ALREADY_EXIST, context)
	}
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("The value is %v.\n", data)
	}

	// Need to run this with goroutine, working on this
	// for index, address := range location.Address {
	// 	externalAddress, _ := handler.externalServicesAdapter.TerminalCreateAddressAdaptor(address)
	// 	if externalAddress["data"] != nil {
	// 		result := externalAddress["data"].(map[string]interface{})
	// 		address_id := result["address_id"].(string)
	// 		location.Address[index].TerminalAddressID = address_id
	// 		terminalLocation := &domain.Location{
	// 			TerminalAddressID: location.Address[index].TerminalAddressID,
	// 			Address: location.Address[index],
	// 			UserID: location.UserID,
	// 		}
	// 		handler.internalServicesAdapter.TerminalUpdateAddressAdaptor(*terminalLocation)
	// 	} else {
	// 		return handler.ComputeErrorResponse(http.StatusBadRequest, externalAddress["message"], context)
	// 	}
	// }
	// Process valid location data
	return handler.ComputeResponseMessage(http.StatusCreated, ADDRESSES_SUBMITTED_SUCCESSFULLY, context)
}


func (handler *HTTPHandler) processAddres(address []domain.Address, location uuid.UUID, result chan <- domain.Location) {
	for index, add := range address {
		externalAddress, err := handler.externalServicesAdapter.TerminalCreateAddressAdaptor(add)
		if err != nil {
			fmt.Println(err)
		}
		resul := externalAddress["data"].(map[string]interface{})
		address_id := resul["address_id"].(string)
		address[index].TerminalAddressID = address_id
		terminalLocation := &domain.Location{
			TerminalAddressID: address[index].TerminalAddressID,
			Address: address[index],
			UserID: location,
		}
		// terminal = append(terminal, externalAddress)
		result <- *terminalLocation
	}
}

func (handler *HTTPHandler) processUpdates(terminalLocation domain.Location) {
	err := handler.internalServicesAdapter.TerminalUpdateAddressAdaptor(terminalLocation)
	if err != nil {
		fmt.Println(err)
	}
}

func (handler *HTTPHandler) worker(_ int, location uuid.UUID, addressChannel <- chan []domain.Address, res chan <- domain.Location) {
	defer wg.Done()
	for address := range addressChannel {
		handler.processAddres(address, location, res)
	}
}

// @Summary Get a address
// @Description get a address
// @Tags Address
// @Accept json
// @Produce json
// @Param addressID path string true "Address ID"
// @Param Authorization header string true "Bearer token"
// @Failure 400 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /addresses/:addressID [get]
func (handler *HTTPHandler) GetAddress(context echo.Context) error {
	user, err := handler.ParseUserID(context)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, err.Error(), context)
	}

	if user.UserType != string(domain.USER) {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}

	ID := context.Param("addressID")
	addressID, err := uuid.Parse(ID)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	location, err := handler.internalServicesAdapter.GetLocationAdaptor(addressID, user.ID);
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusConflict, err.Error(), context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, location, context)
}

// @Summary Get addresses
// @Description get addresses
// @Tags Address
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Failure 400 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /addresses [get]
func (handler *HTTPHandler) GetAddresses(context echo.Context) error {
	user, err := handler.ParseUserID(context)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, err.Error(), context)
	}

	if user.UserType != string(domain.USER) {
		return handler.ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}

	locations, err := handler.internalServicesAdapter.GetLocationsAdaptor(user.ID);
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusConflict, err.Error(), context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, locations, context)
}
