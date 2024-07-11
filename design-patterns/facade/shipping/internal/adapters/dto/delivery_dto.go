package dto

import (
	"github.com/QUDUSKUNLE/shipping/internal/core/model"
)

type DeliveryDTO struct {
	AccountID string  `json:"accountID" binding:"required" validate:"required"`
	PickUpAddress string `json:"pickUpAddress" binding:"required" validate:"required"`
	DeliveryAddress string `json:"deliveryAddress" binding:"required" validate:"required"`
	ProductType model.ProductType `json:"productType" binding:"required" validate:"required"`
}
