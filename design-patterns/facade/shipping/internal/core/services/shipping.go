package services

import (
	"time"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/google/uuid"
	"fmt"
)

func (httpHandler *InternalServicesHandler) NewShippingAdaptor(shippingDto *domain.ShippingDTO) error {
	fmt.Println("Initiate a new shipping")
	systemsHandler := httpHandler.NewInternalServicesFacade()
	newShipping := systemsHandler.shippingService.BuildNewShipping(*shippingDto)
	if err := httpHandler.internal.CreateShippingAdaptor(*newShipping); err != nil {
		return err
	}
	pickUpDTO := domain.PickUpDTO{
		ShippingID: newShipping.ID,
		CarrierID: newShipping.CarrierID,
		Status: string(domain.SCHEDULED),
		PickUpAt: time.Now(),
	}
	if err := httpHandler.NewPickUpAdaptor(pickUpDTO); err != nil {
		return err
	}
	fmt.Println("Shipping created successfully.")
	return nil
}

func (httpHandler *InternalServicesHandler) GetShippingsAdaptor(ID uuid.UUID) ([]domain.Shipping, error) {
	fmt.Println("Initiate a shipping")
	shippings, err := httpHandler.internal.GetShippingsAdaptor(ID, "SCHEDULED")
	if err != nil {
		return []domain.Shipping{}, err
	}
	return shippings, nil
}
