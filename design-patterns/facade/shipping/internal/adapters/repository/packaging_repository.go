package repository

import "github.com/QUDUSKUNLE/shipping/internal/core/domain"

func (database *PostgresRepository) SavePackagingAdaptor(packages []*domain.Packaging) error {
	_ = database.db.AutoMigrate(&domain.Packaging{})
	result := database.db.Create(packages)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
