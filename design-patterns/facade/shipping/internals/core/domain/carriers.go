package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Carrier struct {
		gorm.Model
		ID        uuid.UUID      `json:"id" gorm:"uuid;primaryKey"`
		CreatedAt *time.Time      `json:"created_at"`
		UpdatedAt *time.Time      `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`

		CompanyName    string    `json:"company_name"`
		CompanyAddress Address   `gorm:"embedded" json:"company_address"`
		Contact        Contact   `gorm:"embedded" json:"contact"`
		UserID         uuid.UUID `json:"-"`
		User           *User

		PickUps []PickUp `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	}
	CarrierDto struct {
		CompanyName    string  `json:"company_name" binding:"required" validate:"required,gte=6,lte=1000"`
		CompanyAddress Address `json:"company_address" binding:"required" validate:"required"`
		Contact        Contact `json:"contact" binding:"required" validate:"required,dive,required"`
		UserID         uuid.UUID
	}
)

func (carrier *Carrier) BeforeCreate(scope *gorm.DB) error {
	carrier.ID = uuid.New()
	return nil
}

func (carr *Carrier) BuildNewCarrier(carrierDto CarrierDto) (*Carrier, error) {
	return &Carrier{
		UserID: carrierDto.UserID,
	}, nil
}