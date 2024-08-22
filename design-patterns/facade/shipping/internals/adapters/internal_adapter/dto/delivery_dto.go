package dto

import (
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
)

type DeliveryDto struct {
	AccountID string  `json:"account_id" validate:"required"`
	PickUpAddress string `json:"pick_up_address" validate:"required"`
	DeliveryAddress string `json:"delivery_address" validate:"required"`
	ProductType domain.ProductType `json:"product_type" validate:"required"`
}
