package repository

import (
	"errors"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (database *PostgresRepository) SaveCarrierAdaptor(carrier domain.Carrier) (err error) {
	_ = database.db.AutoMigrate(&domain.Carrier{})
	result := database.db.Create(&domain.Carrier{
		ID:             carrier.ID,
		UserID:         carrier.UserID,
		CompanyName:    carrier.CompanyName,
		CompanyAddress: carrier.CompanyAddress,
		Contact:        carrier.Contact,
	})

	if result.Error != nil {
		return result.Error
	}
	return
}

func (database *PostgresRepository) ReadCarrierAdaptor(ID uuid.UUID) (*domain.Carrier, error) {
	carrier := &domain.Carrier{UserID: ID}
	err := database.db.First(carrier).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Carrier{}, err
	}
	return carrier, nil
}
