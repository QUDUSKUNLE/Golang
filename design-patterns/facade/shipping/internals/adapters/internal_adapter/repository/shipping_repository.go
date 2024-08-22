package repository

import (
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/google/uuid"
)

func (database *PostgresRepository) CreateShippingAdaptor(shipping []*domain.Shipping) error {
	_ = database.db.AutoMigrate(&domain.Shipping{})
	result := database.db.Create(shipping)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (database *PostgresRepository) GetShippingsAdaptor(userID uuid.UUID, status string) ([]*domain.Shipping, error) {
	var shippings []*domain.Shipping
	result := database.db.Preload("Carrier").Find(&shippings, domain.Shipping{UserID: userID}).Limit(10)
	if result.Error != nil {
		return []*domain.Shipping{}, result.Error
	}
	return shippings, nil
}
