package repository

import (
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/google/uuid"
)

func (database *PostgresRepository) CreateShippingAdaptor(shipping domain.Shipping) error {
	result := database.db.Create(&domain.Shipping{
		ID:              shipping.ID,
		UserID:          shipping.UserID,
		Description:     shipping.Description,
		PickUpAddress:   shipping.PickUpAddress,
		DeliveryAddress: shipping.DeliveryAddress,
		ProductType:     shipping.ProductType,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (database *PostgresRepository) GetShippingsAdaptor(userID uuid.UUID, status string) ([]domain.Shipping, error) {
	var shippings []domain.Shipping
	result := database.db.Where(&domain.Shipping{UserID: userID}).Preload("PickUp").Limit(10).Find(&shippings, domain.Shipping{UserID: userID})
	if result.Error != nil {
		return []domain.Shipping{}, result.Error
	}
	return shippings, nil
}
