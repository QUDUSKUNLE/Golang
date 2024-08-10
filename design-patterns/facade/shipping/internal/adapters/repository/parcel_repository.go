package repository

import "github.com/QUDUSKUNLE/shipping/internal/core/domain"

func (database *PostgresRepository) SaveParcelAdaptor(carrier domain.ParcelDto) error {
	result := database.db.Create(&domain.Parcel{
		UserID:    carrier.UserID,
		TerminalParcelID: carrier.TerminalParcelID,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
