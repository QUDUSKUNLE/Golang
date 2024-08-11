package repository

import "github.com/QUDUSKUNLE/shipping/internal/core/domain"

func (database *PostgresRepository) SavePackagingAdaptor(packages []*domain.Packaging) error {
	result := database.db.Create(packages)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
