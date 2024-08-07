package domain

import (
	"time"
	"errors"
	"gorm.io/gorm"
	"github.com/google/uuid"
	"database/sql/driver"
	"encoding/json"
)

type Location struct {
	gorm.Model
	ID              uuid.UUID   `gorm:"primaryKey;->;<-:create" json:"ID"`
	CreatedAt       time.Time   `json:"CreatedAt"`
	UpdatedAt       *time.Time  `json:"UpdatedAt"`
	DeletedAt       *time.Time  `json:"-"`

	Address          Address    `gorm:"embedded" json:"Address"`
	UserID           uuid.UUID  `json:"-"`
	User             *User      `json:"-"`
}

type LocationDTO struct {
	Address Address `json:"Address" binding:"required" validate:"required"`
	UserID  uuid.UUID
}

type Address struct {
	StreetNo   int    `json:"StreetNo" binding:"required,gte=0,let=1000" validate:"required"`
	StreetName string `json:"StreetName" binding:"required,max=50" validate:"required"`
	Province   string `json:"Province" binding:"required,max=50" validate:"required"`
	State      string `json:"State" binding:"required,max=50" validate:"required"`
	Country    Country `json:"Country"`
}

func (a Address) Value() (driver.Value, error) {
	// Serialize the Address struct into a format suitable for storage
	// For example, you might serialize it into a JSON string
	addressJSON, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return string(addressJSON), nil
}

func (a *Address) Scan(value interface{}) error {
	addressJSON, ok := value.(string)
	if !ok {
		return errors.New("unexpected type for address")
	}
	return json.Unmarshal([]byte(addressJSON), a)
}

func (location *Location) BeforeCreate(scope *gorm.DB) error {
	location.ID = uuid.New()
	return nil
}

func (location *Location) BuildNewLocation(locationDto []LocationDTO) ([]*Location) {
	locations := []*Location{}
	for _, location := range locationDto {
		locations = append(locations, &Location{
			UserID: location.UserID,
			Address: location.Address,
		})
	}
	return locations
}
