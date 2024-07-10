package model

import (
	"time"
	"database/sql/driver"
	"errors"
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	gorm.Model
	ID 							uuid.UUID  `gorm:"primaryKey;->;<-:create" json:"id"`
	UserID      		uuid.UUID  `json:"user_id"`
	Description  		string    `json:"description"`
	PickUpAddress 	Address   `gorm:"embedded" json:"pick_up_address"`
	DeliveryAddress Address	  `gorm:"embedded" json:"delivery_address"`
	ProductType  ProductType  `json:"product_type"`
	CreatedAt 	time.Time 		`json:"created_at"`
	UpdatedAt 	*time.Time 		`json:"updated_at"`
	DeletedAt   *time.Time    `gorm:"-:all" json:"-"`
	PickUp      PickUp   			`json:"pick_up"`
}

type ShippingDTO struct {
	Description   	string  	`json:"description" binding:"required" validate:"required,gte=6,lte=1000"`
	PickUpAddress 	Address  	`json:"pick_up_address" binding:"required" validate:"required"`
	DeliveryAddress Address 	`json:"delivery_address" binding:"required" validate:"required"`
	ProductType 		ProductType `json:"product_type" binding:"required" validate:"required"`
}

type Address struct {
	StreetNo   int 		`json:"street_no" binding:"required,gte=0,let=1000" validate:"required"`
	StreetName string `json:"street_name" binding:"required,max=50" validate:"required"`
	Province   string `json:"province" binding:"required,max=50" validate:"required"`
	State      string `json:"state" binding:"required,max=50" validate:"required"`
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
		ID: uuid.New(),
		UserID: userID,
		Description: ship.Description,
		PickUpAddress: ship.PickUpAddress,
		DeliveryAddress: ship.DeliveryAddress,
		ProductType: ship.ProductType,
	}
}
