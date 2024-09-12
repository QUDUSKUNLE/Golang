package integration

import (
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/QUDUSKUNLE/shipping/internals/core/ports"
)

func (externalHandler *ExternalServicesHandler) TerminalCreatePackagingAdaptor(packaging domain.SingleTerminalPackagingDto) (map[string]interface{}, error ){
	var result map[string]interface{}
	serviceHandler := NewExternalServicesFacade()
	buildPackaging := serviceHandler.terminalService.BuildNewTerminalPackaging(packaging)
	result, err := externalHandler.external.TerminalRequest(ports.PACKAGING, ports.POST, buildPackaging)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (externalHandler *ExternalServicesHandler) TerminalCreateAddressAdaptor(address domain.Address) (map[string]interface{}, error ){
	var result map[string]interface{}
	serviceHandler := NewExternalServicesFacade()
	buildAddress := serviceHandler.terminalService.BuildNewTerminalAddress(address)
	result, err := externalHandler.external.TerminalRequest(ports.ADDRESSES, ports.POST, buildAddress)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (externalHandler *ExternalServicesHandler) TerminalGetRatesAdaptor(query domain.TerminalRatesQueryDto) (map[string]interface{}, error ){
	var result map[string]interface{}
	serviceHandler := NewExternalServicesFacade()
	buildQuery := serviceHandler.terminalService.BuildNewTerminalRatesQuery(query)
	result, err := externalHandler.external.TerminalRequest(ports.SHIPMENT, ports.GET, buildQuery)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (externalHandler *ExternalServicesHandler) TerminalCreateParcelAdaptor(parcel domain.SingleTerminalParcelDto) (map[string]interface{}, error ){
	var result map[string]interface{}
	serviceHandler := NewExternalServicesFacade()
	builtParcel := serviceHandler.terminalService.BuildNewTerminalParcel(parcel)
	result, err := externalHandler.external.TerminalRequest(ports.PARCELS, ports.POST, builtParcel)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (externalHandler *ExternalServicesHandler) TerminalCreateShipmentAdaptor(shipment domain.SingleTerminalShipmentDto) (map[string]interface{}, error ){
	var result map[string]interface{}
	serviceHandler := NewExternalServicesFacade()
	builtShipment := serviceHandler.terminalService.BuildNewTerminalShipment(shipment)
	result, err := externalHandler.external.TerminalRequest(ports.SHIPMENTS, ports.POST, builtShipment)
	if err != nil {
		return result, err
	}
	return result, nil
}
