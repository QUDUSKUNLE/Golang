package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Parcel struct {
		gorm.Model
		ID        uuid.UUID      `gorm:"primaryKey;->;<-:create" json:"id"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
		DeletedAt gorm.DeletedAt `json:"deleted_at"`

		UserID           uuid.UUID `json:"-"`
		User             *User     `json:"-"`
		TerminalParcelID string    `json:"terminal_parcel_id"`
	}
	ParcelDto struct {
		UserID           uuid.UUID
		TerminalParcelID string `json:"terminal_parcel_id"`
	}
)

func (parcel *Parcel) BeforeCreate(scope *gorm.DB) error {
	parcel.ID = uuid.New()
	return nil
}

func (parcel *Parcel) BuildNewParcel(packageDto ParcelDto) *ParcelDto {
	return &ParcelDto{
		UserID:           packageDto.UserID,
		TerminalParcelID: packageDto.TerminalParcelID,
	}
}
