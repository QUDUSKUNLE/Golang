package repository

import (
	"errors"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/google/uuid"
)

func (database *PostgresRepository) ReadAddressAdaptor(addressID, userID uuid.UUID) (*domain.Location, error) {
	location := domain.Location{ID: addressID, UserID: userID}
	result := database.db.First(&location)
	if result.RowsAffected == 0 {
		return &domain.Location{}, errors.New("record not found")
	}
	return &location, nil
}

func (database *PostgresRepository) QueryAddressAdaptor(userID uuid.UUID, description string) (*domain.Location, error) {
	location := domain.Location{}
	result := database.db.Find(&location, domain.Location{Description: description, UserID: userID}).Limit(1)
	if result.RowsAffected == 0 {
		return &domain.Location{}, errors.New("record not found")
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
	_ = database.db.AutoMigrate(&domain.Location{})
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

func (database *PostgresRepository) TerminalUpdateAddressAdaptor(location domain.Location) error {
	database.db.Model(&domain.Location{}).Where(&domain.Location{Description: location.Description, UserID: location.UserID}).Updates(map[string]interface{}{"terminal_address_id": location.TerminalAddressID, "address": location.Address})
	return nil
}
