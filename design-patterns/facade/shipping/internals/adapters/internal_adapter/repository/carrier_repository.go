package repository

import (
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
)

func (database *PostgresRepository) SaveCarrierAdaptor(carrier domain.Carrier) error {
	_ = database.db.AutoMigrate(&domain.Carrier{})
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

func (database *PostgresRepository) ReadCarrierAdaptor(ID uuid.UUID) (*domain.Carrier, error) {
	carrier := domain.Carrier{UserID: ID}
	result := database.db.First(&carrier)
	if result.Error != nil {
		return &domain.Carrier{}, result.Error
	}
	return &carrier, nil
}
