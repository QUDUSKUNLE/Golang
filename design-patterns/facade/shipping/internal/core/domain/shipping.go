package domain

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Shipping struct {
	gorm.Model
	ID              uuid.UUID   `gorm:"primaryKey;->;<-:create" json:"ID"`
	CreatedAt       time.Time   `json:"CreatedAt"`
	UpdatedAt       *time.Time  `json:"UpdatedAt"`
	DeletedAt       *time.Time  `json:"-"`

	UserID          uuid.UUID   `json:"-"`
	User            *User       `json:"-"`
	CarrierID       uuid.UUID   `json:"CarrierID"`
	Carrier         *Carrier    `json:"-"`
	Description     string      `gorm:"size=150" json:"Description"`
	PickUpAddress   Address     `gorm:"embedded" json:"PickUpAddress"`
	DeliveryAddress Address     `gorm:"embedded" json:"DeliveryAddress"`
	ProductType     ProductType `json:"ProductType"`
	PickUp          PickUp      `json:"-"`
}

type ShippingDTO struct {
	Description     string      `json:"Description" binding:"required" validate:"required,gte=6,lte=1000"`
	PickUpAddress   Address     `json:"PickUpAddress" binding:"required" validate:"required"`
	DeliveryAddress Address     `json:"DeliveryAddress" binding:"required" validate:"required"`
	ProductType     ProductType `json:"ProductType" binding:"required" validate:"required"`
	CarrierID       uuid.UUID   `json:"CarrierID" binding:"required" validate:"required"`
	UserID          uuid.UUID
}

func (shipping *Shipping) BuildNewShipping(ship ShippingDTO) *Shipping {
	return &Shipping{
		ID:              uuid.New(),
		UserID:          ship.UserID,
		CarrierID:       ship.CarrierID,
		Description:     ship.Description,
		PickUpAddress:   ship.PickUpAddress,
		DeliveryAddress: ship.DeliveryAddress,
		ProductType:     ship.ProductType,
	}
}
