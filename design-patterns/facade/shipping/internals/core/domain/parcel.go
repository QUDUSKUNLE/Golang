package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Parcel struct {
		ID        uuid.UUID      `gorm:"primaryKey;->;<-:create" json:"id"`
		CreatedAt *time.Time      `json:"created_at"`
		UpdatedAt *time.Time      `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `json:"-"`

		TerminalParcelID string    `json:"terminal_parcel_id"`
		UserID           uuid.UUID `json:"-"`
		User             *User     `json:"-"`
	}
	ParcelDto struct {
		UserID           uuid.UUID
		ParcelID []string `json:"parcel_id"`
	}
)

func (parcel *Parcel) BeforeCreate(scope *gorm.DB) error {
	parcel.ID = uuid.New()
	return nil
}

func (parcel *Parcel) BuildNewParcel(parcelDto ParcelDto) []*Parcel {
	parc := []*Parcel{}
	for _, parcelID := range parcelDto.ParcelID {
		parc = append(parc, &Parcel{
			UserID: parcelDto.UserID,
			TerminalParcelID: parcelID,
		})
	}
	return parc
}
