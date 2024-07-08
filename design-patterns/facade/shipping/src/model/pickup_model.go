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
	ID   uuid.UUID `db:"id" json:"id" binding:"required" validate:"required"`
	ShippingID uuid.UUID `db:"shipping_id" json:"shipping_id" binding:"required" validate:"required"`
	CarrierID  uuid.UUID `db:"carrier_id" json:"carrier_id" binding:"required" validate:"required"`
	PickUpAt   time.Time `db:"pick_up_at" json:"pick_up_at" binding:"required" validate:"required"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
	Status     PickUpStatus `db:"status" json:"status" binding:"required" validate:"required"`
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
		CarrierID:  pick.CarrierID,
		PickUpAt:   pick.PickUpAt,
		Status:     PickUpStatus(pick.Status),
	}
}

func (pi *PickUp) BuildUpdatePickUp(pick PickUp) *PickUp {
	return &PickUp{
		ID:         pick.ID,
		ShippingID: pick.ShippingID,
		CarrierID:  pick.CarrierID,
		PickUpAt:   time.Now(),
		Status:     PickUpStatus(pick.Status),
		UpdatedAt:  time.Now(),
	}
}
