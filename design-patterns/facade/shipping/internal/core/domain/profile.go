package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	gorm.Model
	ID        uuid.UUID      `json:"ID" gorm:"uuid;primaryKey"`
	CreatedAt time.Time      `json:"CreatedAt"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	FullName string    `json:"FullName"`
	Address  Address   `gorm:"embedded" json:"CompanyAddress"`
	Contact  Contact   `gorm:"embedded" json:"ContactAddress"`
	UserID   uuid.UUID `json:"CarrierID"`
	User     *User
}
