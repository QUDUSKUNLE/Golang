package repository

import (
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (database *PostgresRepository) InitiatePickUp(pickUp domain.PickUp) error {
	result := database.db.Create(&domain.PickUp{
		ID:         pickUp.ID,
		ShippingID: pickUp.ShippingID,
		UserID:     pickUp.UserID,
		PickUpAt:   pickUp.PickUpAt,
		Status:     pickUp.Status,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (database *PostgresRepository) UpdatePickUp(pickUp domain.PickUp) error {
	query := domain.PickUp{ID: pickUp.ID, UserID: pickUp.UserID}
	database.db.Where(&query).Updates(domain.PickUp{
		PickUpAt:  pickUp.PickUpAt,
		UpdatedAt: pickUp.UpdatedAt,
		Status:    pickUp.Status,
	})
	return nil
}
