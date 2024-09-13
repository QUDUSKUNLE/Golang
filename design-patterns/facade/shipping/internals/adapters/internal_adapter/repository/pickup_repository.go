package repository

import (
	"errors"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/google/uuid"
)

func (database *PostgresRepository) InitiatePickUpAdaptor(pickUp []*domain.PickUp) error {
	_ = database.db.AutoMigrate(&domain.PickUp{})
	result := database.db.Create(pickUp)
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

func (database *PostgresRepository) CarrierPickUps(userID uuid.UUID) ([]*domain.PickUp, error) {
	var pickUps []*domain.PickUp
	carrier := &domain.Carrier{}
	result := database.db.Where("user_id = ?", userID).Find(carrier)
	if result.RowsAffected == 0 {
		return []*domain.PickUp{}, errors.New("record not found")
	}
	result = database.db.Preload("Shipping").Order("created_at desc").Limit(20).Find(&pickUps, domain.PickUp{CarrierID: carrier.ID})
	if result.Error != nil {
		return []*domain.PickUp{}, result.Error
	}
	return pickUps, nil
}

func (database *PostgresRepository) GetPickUp(pickUpID, userID uuid.UUID) (domain.PickUp, error) {
	var pickUp domain.PickUp
	carrier := &domain.Carrier{}
	result := database.db.Where("user_id = ?", userID).Find(carrier)
	if result.RowsAffected == 0 {
		return domain.PickUp{}, errors.New("record not found")
	}
	result = database.db.Preload("Shipping").Find(&pickUp, domain.PickUp{CarrierID: carrier.ID, ID: pickUpID})
	if result.Error != nil {
		return domain.PickUp{}, result.Error
	}
	return pickUp, nil
}
