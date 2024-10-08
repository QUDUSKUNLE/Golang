package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Location struct {
		ID        uuid.UUID       `gorm:"primaryKey;->;<-:create" json:"id"`
		CreatedAt *time.Time      `json:"created_at"`
		UpdatedAt *time.Time      `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`

		TerminalAddressID string    `json:"terminal_address_id"`
		Address           Address   `gorm:"embedded" json:"address"`
		UserID            uuid.UUID `json:"user_id" gorm:"uniqueIndex:idx_description,sort:desc"`
		Description       string    `json:"description" gorm:"uniqueIndex:idx_description,sort:desc"`
		User              *User     `json:"-"`
	}
	LocationDto struct {
		Address []Address `json:"address" validate:"gt=0,dive,required"`
		UserID  uuid.UUID
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
			UserID:            locationDto.UserID,
			Address:           address,
			Description:       address.Description,
			TerminalAddressID: address.TerminalAddressID,
		})
	}
	return locations
}
