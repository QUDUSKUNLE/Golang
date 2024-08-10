package dto

import (
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

type DeliveryDto struct {
	AccountID string  `json:"account_id" binding:"required" validate:"required"`
	PickUpAddress string `json:"pick_up_address" binding:"required" validate:"required"`
	DeliveryAddress string `json:"delivery_address" binding:"required" validate:"required"`
	ProductType domain.ProductType `json:"product_type" binding:"required" validate:"required"`
}
