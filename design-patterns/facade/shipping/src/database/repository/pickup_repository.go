package repository

import (
	"github.com/QUDUSKUNLE/shipping/src/model"
)

func (database *PostgresRepository) QueryCreatePickUp(pickUp model.PickUp) error {
	result := database.Create(&model.PickUp{
		ID: pickUp.ID,
		ShippingID: pickUp.ShippingID,
		UserID: pickUp.UserID,
		PickUpAt: pickUp.PickUpAt,
		Status: pickUp.Status,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (database *PostgresRepository) QueryUpdatePickUp(pickUp model.PickUp) error {
	query := model.PickUp{ID: pickUp.ID, UserID: pickUp.UserID}
	database.Where(&query).Updates(model.PickUp{
		PickUpAt: pickUp.PickUpAt,
		UpdatedAt: pickUp.UpdatedAt,
		Status: pickUp.Status,
	})
	return nil
}
