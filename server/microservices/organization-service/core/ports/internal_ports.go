package ports

import (
	shipping "github.com/QUDUSKUNLE/microservices/organization-service/core/domain"
	"github.com/google/uuid"
)

// Ports that connect internal services of the app
type RepositoryPorts interface {
	CreateShipping(shipping []*shipping.Shipping) error
	GetShippings(shippingID uuid.UUID, status string) ([]*shipping.Shipping, error)
	GetShipping(shippingID, userID uuid.UUID) (*shipping.Shipping, error)
}

// Ports that connect internal services of the app
type ShippingPorts interface {
	CreateShipping(shipping []*shipping.Shipping) error
	GetShippings(shippingID uuid.UUID, status string) ([]*shipping.Shipping, error)
	GetShipping(shippingID, userID uuid.UUID) (*shipping.Shipping, error)
}
