package ports

import (
	// "github.com/google/uuid"
	// domain "github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

type ExternalPorts interface {
	// SaveUser(user domain.UserDTO) error
	// GetUser(UserID uuid.UUID) (*domain.User, error)
	// ResetPassword(Email string) error

	// CarrierUpdatePickUp(pickUp domain.PickUp) error
	// CarrierPickUps() error
	
	// ScheduleShipping(shipping domain.ShippingDTO) error
	// GetShippings(ID uuid.UUID, status string) ([]domain.Shipping, error)
	// UpdateShipping() error

	// Packaging
	TerminalCreatePackaging(packaging interface{}) (interface{}, error)
	TerminalGetPackagingRates() (interface{}, error)

	// Save Addresses
	TerminalCreateAddress(address interface{}) (interface{}, error)
	// UpdateAddress() error
	// GetAddress() error
	// DeleteAddress() error
	// GetAddresses() error

	// ComparePrice() error
	// AddMoneyToWallet() error
	// CheckBalance() error
	// Tracking() error
}

