package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type ProductType string

const (
	Animal      ProductType = "Animal"
	Plant       ProductType = "Plant"
	Appareal    ProductType = "Appareal"
	Book        ProductType = "Book"
	Cosmetics   ProductType = "Cosmetics"
	Electronics ProductType = "Electronics"
	Watery      ProductType = "Watery"
	Ammunition  ProductType = "Ammunition"
	Unknown     ProductType = "Unknown"
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

type Address struct {
	StreetNo   int    `json:"StreetNo" binding:"required,gte=0,let=1000" validate:"required"`
	StreetName string `json:"StreetName" binding:"required,max=50" validate:"required"`
	Province   string `json:"Province" binding:"required,max=50" validate:"required"`
	State      string `json:"State" binding:"required,max=50" validate:"required"`
}

func (a Address) Value() (driver.Value, error) {
	// Serialize the Address struct into a format suitable for storage
	// For example, you might serialize it into a JSON string
	addressJSON, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return string(addressJSON), nil
}

func (a *Address) Scan(value interface{}) error {
	addressJSON, ok := value.(string)
	if !ok {
		return errors.New("unexpected type for address")
	}
	return json.Unmarshal([]byte(addressJSON), a)
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

type Contact struct {
	PhoneNumbers []string `json:"PhoneNumbers"`
	WhatsApps string `json:"WhatsApp" binding:"required,max=50" validate:"required"`
	Twitter  string `json:"Twitter" binding:"required,max=50" validate:"required"`
}

func (a Contact) Value() (driver.Value, error) {
	// Serialize the Address struct into a format suitable for storage
	// For example, you might serialize it into a JSON string
	contactJSON, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return string(contactJSON), nil
}

func (a *Contact) Scan(value interface{}) error {
	contactJSON, ok := value.(string)
	if !ok {
		return errors.New("unexpected type for address")
	}
	return json.Unmarshal([]byte(contactJSON), a)
}
