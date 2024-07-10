package model

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
	ID   uuid.UUID 						`gorm:"primaryKey;->;<-:create" json:"ID"`
	ShippingID uuid.UUID 			`json:"ShippingID"`
	UserID    uuid.UUID 			`json:"CarrierID"`
	PickUpAt   time.Time 			`json:"PickUpAt"`
	CreatedAt 	time.Time 		`json:"CreatedAt"`
	UpdatedAt 	*time.Time 		`json:"UpdatedAt"`
	DeletedAt   *time.Time    `gorm:"-:all" json:"DeletedAt"`
	Status      PickUpStatus 	`json:"Status"`
}

type PickUpDTO struct {
	ShippingID uuid.UUID `json:"shipping_id" binding:"required" validate:"required"`
	CarrierID  uuid.UUID `json:"carrier_id" binding:"required" validate:"required"`
	PickUpAt   time.Time `json:"pick_up_at" binding:"required" validate:"required"`
	Status     string    `json:"status" binding:"required" validate:"required"`
}

func (pickUp *PickUp) BuildNewPickUp(pick PickUpDTO) *PickUp {
	return &PickUp{
		ID:         uuid.New(),
		ShippingID: pick.ShippingID,
		UserID:  pick.CarrierID,
		PickUpAt:   pick.PickUpAt,
		Status:     PickUpStatus(pick.Status),
	}
}

func (pi *PickUp) BuildUpdatePickUp(pick PickUp) *PickUp {
	return &PickUp{
		ID:         pick.ID,
		ShippingID: pick.ShippingID,
		UserID:  pick.UserID,
		PickUpAt:   time.Now(),
		Status:     PickUpStatus(pick.Status),
	}
}
