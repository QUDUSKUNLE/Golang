package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	SCHEDULED PickUpStatus = "SCHEDULED"
	PICKED    PickUpStatus = "PICKED"
	RETURNED  PickUpStatus = "RETURNED"
	REJECTED  PickUpStatus = "REJECTED"
)

type (
	PickUp struct {
		gorm.Model
		ID        uuid.UUID    `gorm:"primaryKey;->;<-:create" json:"id"`
		PickUpAt  time.Time    `json:"pick_up_at"`
		CreatedAt time.Time    `json:"created_at"`
		UpdatedAt *time.Time   `json:"updated_at"`
		DeletedAt *time.Time   `json:"deleted_at"`
		Status    PickUpStatus `json:"status"`

		ShippingID uuid.UUID `json:"-"`
		Shipping   *Shipping `json:"shipping"`
		CarrierID  uuid.UUID `json:"-"`
		Carrier    *Carrier  `json:"-"`
	}
	PickUpDto struct {
		ShippingID uuid.UUID `json:"shipping_id" binding:"required" validate:"required"`
		CarrierID  uuid.UUID `json:"carrier_id" binding:"required" validate:"required"`
		PickUpAt   time.Time `json:"pick_up_at" binding:"required" validate:"required"`
		Status     string    `json:"status" binding:"required" validate:"required"`
	}
	PickUpStatus string
)

func (pickUp *PickUp) BuildNewPickUp(pick PickUpDto) *PickUp {
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
	}
}
