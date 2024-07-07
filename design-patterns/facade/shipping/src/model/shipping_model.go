package model

import (
	"time"

	"github.com/google/uuid"
)

type ProductType string

const (
	Animal ProductType = "Animal"
	Plant ProductType	 = "Plant"
	Appareal ProductType = "Appareal"
	Book ProductType   = "Book"
	Cosmetics ProductType = "Cosmetics"
	Electronics ProductType = "Electronics"
	Watery  ProductType = "Watery"
	Ammunition ProductType = "Ammunition"
	Unknown ProductType = "Unknown"
)

func (product ProductType) PrintProduct() string {
	switch product {
	case Animal:
		return string(Animal)
	case Plant:
		return string(Plant)
	case Appareal:
		return string(Appareal)
	case Book:
		return string(Book)
	case Cosmetics:
		return string(Cosmetics)
	case Electronics:
		return string(Electronics)
	case Watery:
		return string(Watery)
	case Ammunition:
		return string(Ammunition)
	}
	return string(Unknown)
}

type Shipping struct {
	ID 					uuid.UUID
	UserID      uuid.UUID
	Description  string
	PickUpAddress Address
	DeliveryAddress Address
	ProductType ProductType
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ShippingDTO struct {
	Description   	string  	`json:"description" binding:"required" validate:"required"`
	PickUpAddress 	Address  	`json:"pick_up_address" binding:"required" validate:"required"`
	DeliveryAddress Address 	`json:"delivery_address" binding:"required" validate:"required"`
	ProductType 		ProductType `json:"product_type" binding:"required" validate:"required"`
}

type Address struct {
	StreetNo   int 		`json:"street_no" binding:"required" validate:"required"`
	StreetName string `json:"street_name" binding:"required" validate:"required"`
	City       string `json:"city" binding:"required" validate:"required"`
	State      string `json:"state" binding:"required" validate:"required"`
}

func (shipping *Shipping) BuildShipping(userID uuid.UUID, ship ShippingDTO) (*Shipping, error) {
	return &Shipping{
		ID: uuid.New(),
		UserID: userID,
		Description: ship.Description,
		PickUpAddress: ship.PickUpAddress,
		DeliveryAddress: ship.DeliveryAddress,
		ProductType: ship.ProductType,
	}, nil
}
