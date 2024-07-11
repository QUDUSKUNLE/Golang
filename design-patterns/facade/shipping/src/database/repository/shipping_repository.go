package repository

import (
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/google/uuid"
)

func (database *PostgresRepository) QueryCreateShipping(shipping model.Shipping) error {
	result := database.Create(&model.Shipping{
		ID: shipping.ID,
		UserID: shipping.UserID,
		Description: shipping.Description,
		PickUpAddress: shipping.PickUpAddress,
		DeliveryAddress: shipping.DeliveryAddress,
		ProductType: shipping.ProductType,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (database *PostgresRepository) QueryShippings(userID uuid.UUID, status string) ([]model.Shipping, error) {
	var shippings []model.Shipping
	result := database.Where(&model.Shipping{UserID: userID}).Preload("PickUp").Limit(10).Find(&shippings, model.Shipping{UserID: userID});
	if result.Error != nil {
		return []model.Shipping{}, result.Error
	}
	return shippings, nil
}
