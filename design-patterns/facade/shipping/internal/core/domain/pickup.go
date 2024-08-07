package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PickUpStatus string

const (
	SCHEDULED PickUpStatus = "SCHEDULED"
	PICKED    PickUpStatus = "PICKED"
	RETURNED  PickUpStatus = "RETURNED"
	REJECTED  PickUpStatus = "REJECTED"
)

type PickUp struct {
	gorm.Model
	ID         		uuid.UUID    `gorm:"primaryKey;->;<-:create" json:"ID"`
	PickUpAt   		time.Time    `json:"PickUpAt"`
	CreatedAt  		time.Time    `json:"CreatedAt"`
	UpdatedAt  		*time.Time   `json:"UpdatedAt"`
	DeletedAt  		*time.Time   `json:"DeletedAt"`
	Status     		PickUpStatus `json:"Status"`

	ShippingID  	uuid.UUID    `json:"-"`
	Shipping      *Shipping    `json:"Shipping"`
	CarrierID     uuid.UUID    `json:"-"`
	Carrier       *Carrier     `json:"-"`
}

type PickUpDTO struct {
	ShippingID uuid.UUID `json:"ShippingID" binding:"required" validate:"required"`
	CarrierID  uuid.UUID `json:"CarrierID" binding:"required" validate:"required"`
	PickUpAt   time.Time `json:"PickUpAt" binding:"required" validate:"required"`
	Status     string    `json:"Status" binding:"required" validate:"required"`
}

func (pickUp *PickUp) BuildNewPickUp(pick PickUpDTO) *PickUp {
	return &PickUp{
		ID:         	uuid.New(),
		ShippingID: 	pick.ShippingID,
		CarrierID:    pick.CarrierID,
		PickUpAt:   	pick.PickUpAt,
		Status:     	PickUpStatus(pick.Status),
	}
}

func (pi *PickUp) BuildUpdatePickUp(pick PickUp) *PickUp {
	return &PickUp{
		ID:         	pick.ID,
		ShippingID: 	pick.ShippingID,
		CarrierID:    pick.CarrierID,
		PickUpAt:   	time.Now(),
		Status:     	PickUpStatus(pick.Status),
	}
}
