package repository

import (
	"github.com/QUDUSKUNLE/microservices/organization-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	database *gorm.DB
}

// GetShipping implements ports.RepositoryPorts.
func (db *Repository) GetShipping(shippingID uuid.UUID, userID uuid.UUID) (*domain.Shipping, error) {
	panic("unimplemented")
}

// GetShippings implements ports.RepositoryPorts.
func (db *Repository) GetShippings(shippingID uuid.UUID, status string) ([]*domain.Shipping, error) {
	panic("unimplemented")
}

func NewRepository(database *gorm.DB) ports.RepositoryPorts {
	return &Repository{database: database}
}
