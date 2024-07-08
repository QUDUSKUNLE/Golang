package repository

import (
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/src/model"
)

func (database *Database) QueryCreateShipping(shipping model.Shipping) error {
	query := model.Shipping{
		UserID: shipping.ID,
		Description: shipping.Description,
		PickUpAddress: shipping.PickUpAddress,
		DeliveryAddress: shipping.DeliveryAddress,
		ProductType: shipping.ProductType,
	}
	result := database.Create(&query)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (database *Database) QueryShippings(userID uuid.UUID, status string) ([]model.Shipping, error) {
	shippings := []model.Shipping{}
	result := database.Where(map[string]interface{}{"user_id": userID, "status": status}).Find(&shippings);
	if result.Error != nil {
		return shippings, result.Error
	}
	return shippings, nil
}
