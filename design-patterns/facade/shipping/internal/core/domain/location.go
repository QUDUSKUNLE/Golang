package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Location struct {
		gorm.Model
		ID        uuid.UUID       `gorm:"primaryKey;->;<-:create" json:"id"`
		CreatedAt *time.Time      `json:"created_at"`
		UpdatedAt *time.Time      `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`

		TerminalAddressID string    `json:"terminal_address_id"`
		Address           Address   `gorm:"embedded" json:"address"`
		UserID            uuid.UUID `json:"-" gorm:"uniqueIndex:idx_description,sort:desc"`
		Description       string    `json:"description" gorm:"uniqueIndex:idx_description,sort:desc"`
		User              *User     `json:"-"`
	}
	LocationDto struct {
		Address []Address `json:"address" binding:"required" validate:"required,dive,required"`
		UserID  uuid.UUID
	}
	LocationResult struct {
		ExternalAddress map[string]interface{}
		Index           int
	}
)

func (location *Location) BeforeCreate(scope *gorm.DB) error {
	location.ID = uuid.New()
	return nil
}

func (location *Location) BuildNewLocation(locationDto LocationDto) []*Location {
	locations := []*Location{}
	for _, address := range locationDto.Address {
		locations = append(locations, &Location{
			UserID:  locationDto.UserID,
			Address: address,
			Description: address.Description,
		})
	}
	return locations
}
