package repository

import "github.com/QUDUSKUNLE/shipping/internals/core/domain"

func (database *PostgresRepository) SaveParcelAdaptor(parcel []*domain.Parcel) error {
	_ = database.db.AutoMigrate(&domain.Parcel{})
	result := database.db.Create(parcel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
