package repository

import (
	"errors"
	"gorm.io/gorm"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
)

func (database *PostgresRepository) SavePackagingAdaptor(packages []*domain.Packaging) error {
	_ = database.db.AutoMigrate(&domain.Packaging{})
	result := database.db.Create(packages)
	if errors.Is(result.Error, gorm.ErrForeignKeyViolated) {
		return result.Error
	}
	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return result.Error
	}
	return nil
}
