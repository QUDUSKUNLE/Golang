package repository

import (
	"github.com/QUDUSKUNLE/shipping/src/model"
)

func (database *Database) QueryCreatePickUp(pickUp model.PickUp) error {
	query := model.PickUp{
		ID: pickUp.ID,
		ShippingID: pickUp.ShippingID,
		UserID: pickUp.UserID,
		PickUpAt: pickUp.PickUpAt,
		Status: pickUp.Status,
	}
	result := database.Create(&query)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (database *Database) QueryUpdatePickUp(pickUp model.PickUp) error {
	query := model.PickUp{ID: pickUp.ID, UserID: pickUp.UserID}
	database.Where(&query).Updates(model.PickUp{
		PickUpAt: pickUp.PickUpAt,
		UpdatedAt: pickUp.UpdatedAt,
		Status: pickUp.Status,
	})
	return nil
}
