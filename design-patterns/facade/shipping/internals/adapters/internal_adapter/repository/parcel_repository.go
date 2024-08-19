package repository

import (
	"errors"
	"gorm.io/gorm"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
)

func (database *PostgresRepository) SaveParcelAdaptor(parcel []*domain.Parcel) error {
	_ = database.db.AutoMigrate(&domain.Parcel{})
	result := database.db.Create(parcel)
	if errors.Is(result.Error, gorm.ErrForeignKeyViolated) {
		return result.Error
	}
	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return result.Error
	}
	return nil
}
