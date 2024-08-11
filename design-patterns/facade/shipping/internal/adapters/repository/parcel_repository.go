package repository

import "github.com/QUDUSKUNLE/shipping/internal/core/domain"

func (database *PostgresRepository) SaveParcelAdaptor(parcel []*domain.Parcel) error {
	result := database.db.Create(parcel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
