package domain

import (
	"github.com/google/uuid"
	"time"
	"errors"
	"gorm.io/gorm"
)

type (
	Shipping struct {
		ID        uuid.UUID  `gorm:"primaryKey;->;<-:create" json:"id"`
		CreatedAt *time.Time `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
		DeletedAt *time.Time `json:"-"`

		Description        string      `gorm:"size=150" json:"description"`
		TerminalShipmentID string      `json:"terminal_shipment_id"`
		UserID             uuid.UUID   `json:"-"`
		User               *User       `json:"-"`
		CarrierID          uuid.UUID   `json:"carrier_id"`
		Carrier            *Carrier    `json:"-"`
		ProductType        ProductType `json:"Product_type"`
		PickUp             PickUp      `json:"-"`

		PickUpAddressID   uuid.UUID `json:"pick_up_address_id"`
		DeliveryAddressID uuid.UUID `json:"delivery_address_id"`
		Address           *Address  `json:"-"`
	}
	ShippingDto struct {
		Shipments []SingleShippingDto `json:"shipments" binding:"required" validate:"required,dive,required"`
	}
	SingleShippingDto struct {
		Description       string      `json:"description" validate:"required,gte=6,lte=100"`
		PickUpAddressID   uuid.UUID   `json:"pick_up_address_id" validate:"required"`
		DeliveryAddressID uuid.UUID   `json:"delivery_address_id" validate:"required"`
		PickUpAddress     Address     `json:"pick_up_address"`
		DeliveryAddress   Address     `json:"delivery_address"`
		ProductType       ProductType `json:"product_type" validate:"required"`
		CarrierID         uuid.UUID   `json:"carrier_id" validate:"required"`
		UserID            uuid.UUID
	}
)

func (shipping *Shipping) BeforeCreate(scope *gorm.DB) (err error){
	if shipping.PickUpAddressID == shipping.DeliveryAddressID {
		err = errors.New("pick_up_address_id can`t be same as delivery_address")
	}
	return
}

func (shipping *Shipping) BuildNewShipping(ship ShippingDto) []*Shipping {
	shipments := []*Shipping{}
	for _, shipment := range ship.Shipments {
		shipments = append(shipments, &Shipping{
			ID:                uuid.New(),
			UserID:            shipment.UserID,
			CarrierID:         shipment.CarrierID,
			Description:       shipment.Description,
			PickUpAddressID:   shipment.PickUpAddressID,
			DeliveryAddressID: shipment.DeliveryAddressID,
			ProductType:       shipment.ProductType,
		})
	}
	return shipments
}

func (shipping *Shipping) BuildPickUp(shipments []*Shipping) PickUpDto {
	shippings := PickUpDto{}
	for _, shipment := range shipments {
		shippings.PickUps = append(shippings.PickUps, SinglePickUpDto{
			CarrierID:  shipment.CarrierID,
			ShippingID: shipment.ID,
			PickUpAt:   time.Now(),
			Status:     string(SCHEDULED),
		})
	}
	return shippings
}
