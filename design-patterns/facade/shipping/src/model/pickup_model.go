package model

import (
	"fmt"
	"time"
	"github.com/google/uuid"
)

type PickUpStatus string

const (
	SCHEDULED PickUpStatus = "SCHEDULED"
	PICKED  PickUpStatus = "PICKED"
	RETURNED PickUpStatus = "RETURNED"
	REJECTED PickUpStatus = "REJECTED"
)

type PickUp struct {
	ID 		   		uuid.UUID  		`db:"id"`
	ShippingID	uuid.UUID  		`db:"shipping_id"`
	CarrierID		uuid.UUID  		`db:"carrier_id"`
	PickUpAt    time.Time  		`db:"pick_up_at"`
	CreatedAt 	time.Time  		`db:"created_at"`
	UpdatedAt 	time.Time  		`db:"updated_at"`
	Status      PickUpStatus  `db:"status"`
}

type PickUpDTO struct {
	ShippingID uuid.UUID `json:"shipping_id" binding:"required" validate:"required"`
	CarrierID uuid.UUID `json:"carrier_id" binding:"required" validate:"required"`
	PickUpAt time.Time `json:"pick_up_at" binding:"required" validate:"required"`
	Status  string  `json:"status" binding:"required" validate:"required"`
}

func (pickUp *PickUp) BuildNewPickUp(pick PickUpDTO) *PickUp {
	return &PickUp{
		ID: uuid.New(),
		ShippingID: pick.ShippingID,
		CarrierID: pick.CarrierID,
		PickUpAt: pick.PickUpAt,
		Status: PickUpStatus(pick.Status),
	}
}

func (pickUp *PickUp) BuildUpdatePickUp(pick PickUpDTO) *PickUp {
	return &PickUp{
		ID: pickUp.ID,
		ShippingID: pick.ShippingID,
		CarrierID: pick.CarrierID,
		PickUpAt: pick.PickUpAt,
		Status: PickUpStatus(pick.Status),
	}
}

func (pickUp *PickUp) NewLedger(pick PickUp) (shippingID string, err error) {
	fmt.Printf("Make pick up ledger entry for accountID %s with productType %s.\n", pick.ID, pick.Status)
	return "1", nil
}

func (ledger *PickUp) UpdateLedger(update PickUp) (shippingID string, err error) {
	fmt.Printf("Make pick up ledger entry for accountID %s with productType %s.\n", update.ID, update.Status)
	return "1", nil
}
