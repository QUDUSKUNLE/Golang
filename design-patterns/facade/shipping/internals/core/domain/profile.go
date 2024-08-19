package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	ID        uuid.UUID      `json:"id" gorm:"uuid;primaryKey"`
	CreatedAt *time.Time      `json:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`

	FullName string    `json:"full_name"`
	Address  Address   `gorm:"embedded" json:"company_address"`
	Contact  Contact   `gorm:"embedded" json:"contact_address"`
	UserID   uuid.UUID `json:"carrier_id"`
	User     *User
}
