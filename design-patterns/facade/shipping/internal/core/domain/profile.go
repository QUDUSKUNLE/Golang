package domain

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Profile struct {
	gorm.Model
	ID        uuid.UUID  `json:"ID" gorm:"uuid;primaryKey"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt *time.Time `json:"UpdatedAt,omitempty"`
	DeletedAt  		*time.Time   `json:"DeletedAt"`

	FullName 	string   `json:"FullName"`
	Address 	Address  `gorm:"embedded" json:"CompanyAddress"`
	Contact 	Contact  `gorm:"embedded" json:"ContactAddress"`
	UserID    uuid.UUID  `json:"CarrierID"`
	User      *User
}
