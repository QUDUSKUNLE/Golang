package services

import (
	"time"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/google/uuid"
	"fmt"
)

type ShippingAdaptor struct {
	shippingService *domain.Shipping
}

func (httpHandler *ServicesHandler) NewShippingAdaptor(shippingDto *domain.ShippingDTO) error {
	fmt.Println("Initiate a new shipping")
	adaptor := &ShippingAdaptor{
		shippingService: &domain.Shipping{},
	}
	newShipping := adaptor.shippingService.BuildNewShipping(*shippingDto)
	if err := httpHandler.Internal.CreateShippingAdaptor(*newShipping); err != nil {
		return err
	}
	pickUpDTO := domain.PickUpDTO{
		ShippingID: newShipping.ID,
		CarrierID: newShipping.UserID,
		Status: string(domain.SCHEDULED),
		PickUpAt: time.Now(),
	}
	if err := httpHandler.NewPickUpAdaptor(pickUpDTO); err != nil {
		return err
	}
	fmt.Println("Shipping created successfully.")
	return nil
}

func (httpHandler *ServicesHandler) GetShippingsAdaptor(ID uuid.UUID) ([]domain.Shipping, error) {
	fmt.Println("Initiate a shipping")
	shippings, err := httpHandler.Internal.GetShippingsAdaptor(ID, "SCHEDULED")
	if err != nil {
		return []domain.Shipping{}, err
	}
	return shippings, nil
}
