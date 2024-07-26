package repository

import (
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (database *PostgresRepository) InitiatePickUpAdaptor(pickUp domain.PickUp) error {
	result := database.db.Create(&domain.PickUp{
		ID:         pickUp.ID,
		ShippingID: pickUp.ShippingID,
		CarrierID:     pickUp.CarrierID,
		PickUpAt:   pickUp.PickUpAt,
		Status:     pickUp.Status,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (database *PostgresRepository) UpdatePickUpAdaptor(pickUp domain.PickUp) error {
	query := domain.PickUp{ID: pickUp.ID, CarrierID: pickUp.CarrierID}
	database.db.Where(&query).Updates(domain.PickUp{
		PickUpAt:  pickUp.PickUpAt,
		UpdatedAt: pickUp.UpdatedAt,
		Status:    pickUp.Status,
	})
	return nil
}
