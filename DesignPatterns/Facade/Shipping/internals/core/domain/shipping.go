package domain

import (
	"github.com/google/uuid"
	"time"
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

		PickUpAddressID   uuid.UUID `json:"pickup_address_id"`
		DeliveryAddressID uuid.UUID `json:"delivery_address_id"`
		Address           *Address  `json:"-"`
	}
	ShippingDto struct {
		Shipments []SingleShippingDto `json:"shipments" validate:"gt=0,dive,required"`
	}
	SingleShippingDto struct {
		Description       string      `json:"description" validate:"required,gte=6,lte=100"`
		PickUpAddressID   uuid.UUID   `json:"pickup_address_id" validate:"uuid"`
		// Need to work PickUpAddressID should not be equal to DeliveryAddressID
		DeliveryAddressID uuid.UUID   `json:"delivery_address_id" validate:"uuid,nefield=PickUpAddressID,required"`
		PickUpAddress     Address     `json:"pickup_address" validate:"required_without=PickUpAddressID"`
		DeliveryAddress   Address     `json:"delivery_address" validate:"required_without=DeliveryAddressID"`
		TerminalShipmentID  string    `json:"terminal_shipment_id"`
		ProductType       ProductType `json:"product_type" validate:"required"`
		CarrierID         uuid.UUID   `json:"carrier_id" validate:"uuid,required"`
		UserID            uuid.UUID
	}
)

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
			TerminalShipmentID: shipment.TerminalShipmentID,
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
			Status:     SCHEDULED,
		})
	}
	return shippings
}
