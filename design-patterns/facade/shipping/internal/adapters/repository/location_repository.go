package repository

import (
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/google/uuid"
)

func (database *PostgresRepository) ReadAddressAdaptor(addressID, userID uuid.UUID) (*domain.Location, error) {
	location := domain.Location{ID: addressID, UserID: userID}
	result := database.db.First(&location)
	if result.Error != nil {
		return &domain.Location{}, result.Error
	}
	return &location, nil
}

func (database *PostgresRepository) ReadAddressesAdaptor(userID uuid.UUID) ([]domain.Location, error) {
	var locations []domain.Location
	result := database.db.Find(&locations, domain.Location{UserID: userID}).Limit(10);
	if result.Error != nil {
		return []domain.Location{}, nil
	}
	return locations, nil
}

func (database *PostgresRepository) SaveAddressAdaptor(locations []*domain.Location) error {
	result := database.db.Create(locations)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (database *PostgresRepository) UpdateAddressAdaptor(addressID uuid.UUID, location domain.Location) error {
	database.db.Where(&domain.Location{ID: addressID, UserID: location.UserID}).Updates(domain.Location{Address: location.Address, UpdatedAt: location.UpdatedAt})
	return nil
}

func (database *PostgresRepository) DeleteAddressAdaptor(addressID uuid.UUID) error {
	database.db.Delete(&domain.Location{}, addressID)
	return nil
}
