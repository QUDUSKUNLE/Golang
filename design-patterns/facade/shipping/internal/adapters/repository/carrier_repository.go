package repository

import (
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (database *PostgresRepository) SaveCarrierAdaptor(carrier domain.Carrier) error {
	result := database.db.Create(&domain.Carrier{
		ID:         carrier.ID,
		UserID:    carrier.UserID,
		CompanyName: carrier.CompanyName,
		CompanyAddress: carrier.CompanyAddress,
		Contact: carrier.Contact,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
