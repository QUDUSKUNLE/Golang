package db

import (
	"errors"

	"github.com/QUDUSKUNLE/microservices/organization-service/core/domain"
	"github.com/google/uuid"
)

func (db *Repository) CreateShipping(shipping []*domain.Shipping) error {
	_ = db.database.AutoMigrate(&domain.Shipping{})
	result := db.database.Create(shipping)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (database *PostgresRepository) GetShippingsAdaptor(userID uuid.UUID, status string) ([]*domain.Shipping, error) {
	var shippings []*domain.Shipping
	result := database.db.Order("created_at desc").Limit(20).Find(&shippings, domain.Shipping{UserID: userID})
	if result.Error != nil {
		return []*domain.Shipping{}, result.Error
	}
	return shippings, nil
}

func (database *PostgresRepository) GetShippingAdaptor(shipmentID, userID uuid.UUID) (*domain.Shipping, error) {
	shipping := &domain.Shipping{}
	result := database.db.Limit(1).Find(shipping, domain.Shipping{ID: shipmentID, UserID: userID})
	if result.RowsAffected == 0 {
		return &domain.Shipping{}, errors.New("record not found")
	}
	return shipping, nil
}
