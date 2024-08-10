package services

import (
	"fmt"

	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (httpHandler *ExternalServicesHandler) TerminalPackagingAdaptor(packaging domain.PackagingDto) (interface{}, error ){
	fmt.Println("Initiate a new packaging")
	var result interface{}
	serviceHandler := httpHandler.NewExternalServicesFacade()
	buildPackaging := serviceHandler.terminalService.BuildNewPackaging(packaging)

	result, err := httpHandler.external.TerminalCreatePackaging(buildPackaging)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (httpHandler *ExternalServicesHandler) TerminalAddressAdaptor(address domain.Address) (map[string]interface{}, error ){
	fmt.Println("Initiate a new address")
	var result map[string]interface{}
	serviceHandler := httpHandler.NewExternalServicesFacade()
	buildAddress := serviceHandler.terminalService.BuildNewAddress(address)

	result, err := httpHandler.external.TerminalCreateAddress(buildAddress)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (httpHandler *ExternalServicesHandler) TerminalGetRatesAdaptor() (interface{}, error ){
	fmt.Println("Initiate a new rates")
	var result interface{}
	result, err := httpHandler.external.TerminalGetPackagingRates()
	if err != nil {
		return result, err
	}
	return result, nil
}
