package ports

import (
	"github.com/google/uuid"
	domain "github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

type ServicePorts interface {
	SaveUserAdaptor(user domain.User) error
	ReadUser(UserID uuid.UUID) (*domain.User, error)
	ReadUserByEmail(Email string) (*domain.User, error)
	SaveUser(user domain.User) error

	InitiatePickUp(pickUp domain.PickUp) error
	UpdatePickUp(pickUp domain.PickUp) error

	CreateShipping(shipping domain.Shipping) error
	GetShippings(ID uuid.UUID, status string) ([]domain.Shipping, error)
}

