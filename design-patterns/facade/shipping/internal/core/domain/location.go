package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Location struct {
	gorm.Model
	ID        uuid.UUID      `gorm:"primaryKey;->;<-:create" json:"ID"`
	CreatedAt time.Time      `json:"CreatedAt"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	TerminalAddressID *string   `json:"TerminalAddressID"`
	Address           Address   `gorm:"embedded" json:"Address"`
	UserID            uuid.UUID `json:"-"`
	User              *User     `json:"-"`
}

type LocationDTO struct {
	Address []Address `json:"Address" binding:"required" validate:"required"`
	UserID  uuid.UUID
}

func (location *Location) BeforeCreate(scope *gorm.DB) error {
	location.ID = uuid.New()
	return nil
}

func (location *Location) BuildNewLocation(locationDto LocationDTO) []*Location {
	locations := []*Location{}
	for _, address := range locationDto.Address {
		locations = append(locations, &Location{
			UserID:  locationDto.UserID,
			Address: address,
		})
	}
	return locations
}
