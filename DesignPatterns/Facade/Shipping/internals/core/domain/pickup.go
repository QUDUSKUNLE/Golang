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
		ID        uuid.UUID    `gorm:"primaryKey;->;<-:create" json:"id"`
		PickUpAt  time.Time    `json:"pick_up_at"`
		CreatedAt time.Time    `json:"created_at"`
		UpdatedAt *time.Time   `json:"updated_at"`
		DeletedAt *time.Time   `json:"-"`
		Status    PickUpStatus `json:"status"`

		ShippingID uuid.UUID `json:"-"`
		Shipping   *Shipping `json:"shipping"`
		CarrierID  uuid.UUID `json:"-"`
		Carrier    *Carrier  `json:"-"`
	}
	PickUpDto struct {
		PickUps []SinglePickUpDto `json:"pick_ups" validate:"gt=0,dive,required"`
	}
	SinglePickUpDto struct {
		ID         uuid.UUID    `json:"id"`
		ShippingID uuid.UUID    `json:"shipping_id" validate:"uuid,required"`
		CarrierID  uuid.UUID    `json:"carrier_id" validate:"uuid,required"`
		PickUpAt   time.Time    `json:"pick_up_at" validate:"required"`
		Status     PickUpStatus `json:"status" validate:"required"`
	}
	PickUpStatus string
)

func (pickUp *PickUp) BeforeCreate(scope *gorm.DB) error {
	pickUp.ID = uuid.New()
	return nil
}

func (pickUp *PickUp) BuildNewPickUp(pick PickUpDto) []*PickUp {
	pickUps := []*PickUp{}
	for _, p := range pick.PickUps {
		pickUps = append(pickUps, &PickUp{
			ShippingID: p.ShippingID,
			CarrierID:  p.CarrierID,
			PickUpAt:   p.PickUpAt,
			Status:     PickUpStatus(p.Status),
		})
	}
	return pickUps
}

func (pickUp *PickUp) BuildUpdatePickUp(pick PickUpDto) []*PickUp {
	pickUps := []*PickUp{}
	for _, p := range pick.PickUps {
		pickUps = append(pickUps, &PickUp{
			ID:         p.ID,
			ShippingID: p.ShippingID,
			CarrierID:  p.CarrierID,
			PickUpAt:   time.Now(),
			Status:     PickUpStatus(p.Status),
		})
	}
	return pickUps
}
