package ports

import (
	"github.com/google/uuid"
	domain "github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

type RepositoryPorts interface {
	SaveUserAdaptor(user domain.User) error
	ReadUserAdaptor(UserID uuid.UUID) (*domain.User, error)
	ReadUserByEmailAdaptor(Email string) (*domain.User, error)
	// SaveUser(user domain.User) error

	InitiatePickUpAdaptor(pickUp domain.PickUp) error
	UpdatePickUpAdaptor(pickUp domain.PickUp) error

	CreateShippingAdaptor(shipping domain.Shipping) error
	GetShippingsAdaptor(ID uuid.UUID, status string) ([]domain.Shipping, error)
}

