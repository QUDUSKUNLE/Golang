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
	UserID          uuid.UUID   `json:"UserID"`
	Description     string      `gorm:"size=150" json:"Description"`
	PickUpAddress   Address     `gorm:"embedded" json:"PickUpAddress"`
	DeliveryAddress Address     `gorm:"embedded" json:"DeliveryAddress"`
	ProductType     ProductType `json:"ProductType"`
	CreatedAt       time.Time   `json:"CreatedAt"`
	UpdatedAt       *time.Time  `json:"UpdatedAt"`
	DeletedAt       *time.Time  `gorm:"-:all" json:"DeletedAt"`
	PickUp          PickUp      `json:"PickUp"`
}

type ShippingDTO struct {
	Description     string      `json:"Description" binding:"required" validate:"required,gte=6,lte=1000"`
	PickUpAddress   Address     `json:"PickUpAddress" binding:"required" validate:"required"`
	DeliveryAddress Address     `json:"DeliveryAddress" binding:"required" validate:"required"`
	ProductType     ProductType `json:"ProductType" binding:"required" validate:"required"`
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

func (shipping *Shipping) BuildNewShipping(userID uuid.UUID, ship ShippingDTO) *Shipping {
	return &Shipping{
		ID:              uuid.New(),
		UserID:          userID,
		Description:     ship.Description,
		PickUpAddress:   ship.PickUpAddress,
		DeliveryAddress: ship.DeliveryAddress,
		ProductType:     ship.ProductType,
	}
}
