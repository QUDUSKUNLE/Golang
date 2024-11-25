package services

import (
	"github.com/QUDUSKUNLE/microservices/shipping-service/core/domain"
	"github.com/google/uuid"
)

// GetByEmail implements v1.RepoInterface.
func (use *ServicesHandler) CreateShipping(shipping []*domain.Shipping) error {
	return use.internal.CreateShipping(shipping)
}

// GetShipping implements ports.ShippingPorts.
func (use *ServicesHandler) GetShipping(shippingID uuid.UUID, userID uuid.UUID) (*domain.Shipping, error) {
	panic("unimplemented")
}

// GetShippings implements ports.ShippingPorts.
func (use *ServicesHandler) GetShippings(shippingID uuid.UUID, status string) ([]*domain.Shipping, error) {
	panic("unimplemented")
}
