package ports

import (
	"github.com/google/uuid"
	domain "github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

type ExternalPorts interface {
	SaveUser(user domain.UserDTO) error
	GetUser(UserID uuid.UUID) (*domain.User, error)

	PickUp(pickUp domain.PickUpDTO) error
	UpdatePickUp(pickUp domain.PickUp) error

	Shipping(shipping domain.ShippingDTO) error
	GetShippings(ID uuid.UUID, status string) ([]domain.Shipping, error)
}

