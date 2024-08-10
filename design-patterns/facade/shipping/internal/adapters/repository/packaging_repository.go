package repository

import "github.com/QUDUSKUNLE/shipping/internal/core/domain"

func (database *PostgresRepository) SavePackagingAdaptor(pack domain.PackagingDto) error {
	result := database.db.Create(&domain.Packaging{
		UserID:    pack.UserID,
		TerminalPackagingID: pack.TerminalPackagingID,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
