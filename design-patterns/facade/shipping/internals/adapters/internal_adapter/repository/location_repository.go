package repository

import (
	"errors"

	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (database *PostgresRepository) ReadAddressAdaptor(addressID, userID uuid.UUID) (*domain.Location, error) {
	location := &domain.Location{ID: addressID, UserID: userID}
	err := database.db.First(location).Error;
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Location{}, err
	}
	return location, nil
}

func (database *PostgresRepository) QueryAddressAdaptor(userID uuid.UUID, description string) (*domain.Location, error) {
	location := &domain.Location{}
	result := database.db.Find(location, domain.Location{Description: description, UserID: userID}).Limit(1)
	if result.RowsAffected == 0 {
		return &domain.Location{}, errors.New("record not found")
	}
	return location, nil
}

func (database *PostgresRepository) ReadAddressesAdaptor(userID uuid.UUID) ([]*domain.Location, error) {
	var locations []*domain.Location
	result := database.db.Find(&locations, domain.Location{UserID: userID}).Limit(10);
	if result.Error != nil {
		return []*domain.Location{}, result.Error
	}
	return locations, nil
}

func (database *PostgresRepository) SaveAddressAdaptor(locations []*domain.Location) (err error) {
	_ = database.db.AutoMigrate(&domain.Location{})
	result := database.db.Create(locations)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (database *PostgresRepository) UpdateAddressAdaptor(addressID uuid.UUID, location domain.Location) (err error) {
	database.db.Where(&domain.Location{ID: addressID, UserID: location.UserID}).Updates(domain.Location{Address: location.Address, UpdatedAt: location.UpdatedAt})
	return
}

func (database *PostgresRepository) DeleteAddressAdaptor(addressID uuid.UUID) (err error) {
	database.db.Delete(&domain.Location{}, addressID)
	return
}

func (database *PostgresRepository) TerminalUpdateAddressAdaptor(location domain.Location) (err error) {
	database.db.Model(&domain.Location{Description: location.Description, UserID: location.UserID}).Where(&domain.Location{Description: location.Description, UserID: location.UserID}).Updates(map[string]interface{}{"terminal_address_id": location.TerminalAddressID, "address": location.Address})
	return
}
