package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Shipping struct {
		gorm.Model
		ID        uuid.UUID  `gorm:"primaryKey;->;<-:create" json:"id"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`

		UserID          uuid.UUID   `json:"-"`
		User            *User       `json:"-"`
		CarrierID       uuid.UUID   `json:"carrier_id"`
		Carrier         *Carrier    `json:"-"`
		Description     string      `gorm:"size=150" json:"description"`
		PickUpAddress   Address     `gorm:"embedded" json:"pick_up_address"`
		DeliveryAddress Address     `gorm:"embedded" json:"delivery_address"`
		ProductType     ProductType `json:"Product_type"`
		PickUp          PickUp      `json:"-"`
	}
	ShippingDto struct {
		Description     string      `json:"description" binding:"required" validate:"required,gte=6,lte=1000"`
		PickUpAddress   Address     `json:"pick_up_address" binding:"required" validate:"required,dive,required"`
		DeliveryAddress Address     `json:"delivery_address" binding:"required" validate:"required,dive,required"`
		ProductType     ProductType `json:"product_type" binding:"required" validate:"required"`
		CarrierID       uuid.UUID   `json:"carrier_id" binding:"required" validate:"required"`
		UserID          uuid.UUID
	}
)

func (shipping *Shipping) BuildNewShipping(ship ShippingDto) *Shipping {
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
