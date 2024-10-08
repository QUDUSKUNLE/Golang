package ports

import (
	"github.com/google/uuid"
	domain "github.com/QUDUSKUNLE/shipping/internals/core/domain"
)

// Ports that connect internal services of the app
type RepositoryPorts interface {
	SaveUserAdaptor(user domain.User) (error)
	ReadUserAdaptor(userID uuid.UUID) (*domain.User, error)
	ReadUserByEmailAdaptor(email string) (*domain.User, error)
	// SaveUser(user domain.User) error

	InitiatePickUpAdaptor(pickUp []*domain.PickUp) error
	UpdatePickUpAdaptor(pickUp domain.PickUp) error
	SaveCarrierAdaptor(carrier domain.Carrier) error

	CreateShippingAdaptor(shipping []*domain.Shipping) error
	GetShippingsAdaptor(shippingID uuid.UUID, status string) ([]*domain.Shipping, error)

	CarrierPickUps(carrierID uuid.UUID) ([]*domain.PickUp, error)
	GetPickUp(pickUpID, userID uuid.UUID) (domain.PickUp, error)

	// Addresses Ports
	ReadAddressAdaptor(addressID, userID uuid.UUID) (*domain.Location, error)
	ReadAddressesAdaptor(userID uuid.UUID) ([]*domain.Location, error)
	ReadMultipleAddressesAdaptor(locationIDs []uuid.UUID, userID uuid.UUID) ([]*domain.Location, error)
	SaveAddressAdaptor(location []*domain.Location) error
	UpdateAddressAdaptor(addressID uuid.UUID, location domain.Location) error
	DeleteAddressAdaptor(addressID uuid.UUID) error
	QueryAddressAdaptor(userID uuid.UUID, description string) (*domain.Location, error)
	TerminalUpdateAddressAdaptor(location domain.Location) error
	// Packaging Ports
	SavePackagingAdaptor(pack []*domain.Packaging) error
	// Parcel Ports
	SaveParcelAdaptor(parcel []*domain.Parcel) error
	GetParcelsAdaptor(userID uuid.UUID) ([]*domain.Parcel, error)
}
