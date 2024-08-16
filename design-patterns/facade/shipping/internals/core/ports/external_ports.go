package ports

// Ports that connect external services or third party services
type ExternalPorts interface {
	// Packaging
	TerminalCreatePackaging(packaging interface{}) (map[string]interface{}, error)
	TerminalUpdatePackaging(packaging interface{}) (map[string]interface{}, error)
	TerminalDeletePackaging(packaging interface{}) (map[string]interface{}, error)
	TerminalGetPackaging(packaging interface{}) (map[string]interface{}, error)
	// Get Rates
	TerminalGetRates(query string) (map[string]interface{}, error)
	// Save Addresses
	TerminalCreateAddress(address interface{}) (map[string]interface{}, error)
	// Save Parcel
	TerminalCreateParcel(parcel interface{}) (map[string]interface{}, error)
	TerminalGetParcels() (map[string]interface{}, error)
	TerminalGetParcel(parcel interface{}) (map[string]interface{}, error)
	TerminalUpdateParcel(parcel interface{}) (map[string]interface{}, error)
	TerminalDeleteParcel(parcel interface{}) (map[string]interface{}, error)
	// Create Shipment
	TerminalCreateShipment(shipment interface{}) (map[string]interface{}, error)
	TerminalTrackShipment(shipment interface{}) (map[string]interface{}, error)
	TerminalCancelShipment(shipment interface{}) (map[string]interface{}, error)
	TerminalDeleteShipment(shipment interface{}) (map[string]interface{}, error)
	TerminalGetShipment(shipment interface{}) (map[string]interface{}, error)
	// UpdateAddress() error
	// GetAddress() error
	// DeleteAddress() error
	// GetAddresses() error

	// ComparePrice() error
	// AddMoneyToWallet() error
	// CheckBalance() error
	// Tracking() error
}

