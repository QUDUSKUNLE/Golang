package dto

import (
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

type DeliveryDTO struct {
	AccountID string  `json:"accountID" binding:"required" validate:"required"`
	PickUpAddress string `json:"pickUpAddress" binding:"required" validate:"required"`
	DeliveryAddress string `json:"deliveryAddress" binding:"required" validate:"required"`
	ProductType domain.ProductType `json:"productType" binding:"required" validate:"required"`
}
