package repository

import (
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
)

func (database *PostgresRepository) SaveParcelAdaptor(parcel []*domain.Parcel) error {
	_ = database.db.AutoMigrate(&domain.Parcel{})
	result := database.db.Create(parcel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (database *PostgresRepository) GetParcelsAdaptor(userID uuid.UUID) ([]*domain.Parcel, error) {
	var parcels []*domain.Parcel
	result := database.db.Where("user_id = ?", userID).Find(&parcels).Limit(10)
	if result.Error != nil {
		return []*domain.Parcel{}, result.Error
	}
	return parcels, nil
}
