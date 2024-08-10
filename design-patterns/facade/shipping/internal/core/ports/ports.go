package ports

// Ports that connect external services
type ExternalPorts interface {
	// Packaging
	TerminalCreatePackaging(packaging interface{}) (map[string]interface{}, error)
	TerminalGetRates() (map[string]interface{}, error)
	// Save Addresses
	TerminalCreateAddress(address interface{}) (map[string]interface{}, error)

	// Save Parcel
	TerminalCreateParcel(parcel interface{}) (map[string]interface{}, error)
	// UpdateAddress() error
	// GetAddress() error
	// DeleteAddress() error
	// GetAddresses() error

	// ComparePrice() error
	// AddMoneyToWallet() error
	// CheckBalance() error
	// Tracking() error
}

