package domain

import (
	"time"

	// "github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/google/uuid"
)

type (
	Shipping struct {
		ID        uuid.UUID  `gorm:"primaryKey;->;<-:create" json:"id"`
		CreatedAt *time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
		DeletedAt *time.Time `json:"-"`

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
	Shipments []SingleShippingDto `json:"shipments" binding:"required" validate:"required,dive,required"`
	}
	SingleShippingDto struct {
		Description     string      `json:"description" validate:"required,gte=6,lte=100"`
		PickUpAddress   Address     `json:"pick_up_address" validate:"required"`
		DeliveryAddress Address     `json:"delivery_address" validate:"required"`
		ProductType     ProductType `json:"product_type" validate:"required"`
		CarrierID       uuid.UUID   `json:"carrier_id" validate:"required"`
		UserID          uuid.UUID
	}
)

func (shipping *Shipping) BuildNewShipping(ship ShippingDto) []*Shipping {
	shipments := []*Shipping{}
	for _, shipment := range ship.Shipments {
		shipments = append(shipments, &Shipping{
			ID: 						 uuid.New(),
			UserID:       	 shipment.UserID,
			CarrierID:       shipment.CarrierID,
			Description:     shipment.Description,
			PickUpAddress:   shipment.PickUpAddress,
			DeliveryAddress: shipment.DeliveryAddress,
			ProductType:     shipment.ProductType,
		})
	}
	return shipments
}

func (shipping *Shipping) BuildPickUp(shipments []*Shipping) PickUpDto {
	shippings := PickUpDto{}
	for _, shipment := range shipments {
		shippings.PickUps = append(shippings.PickUps, SinglePickUpDto{
			CarrierID: shipment.CarrierID,
			ShippingID: shipment.ID,
			PickUpAt: time.Now(),
			Status: string(SCHEDULED),
		})
	}
	return shippings
}
