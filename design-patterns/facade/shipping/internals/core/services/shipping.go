package services

import (
	"time"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/google/uuid"
)

func (internalHandler *InternalServicesHandler) NewShippingAdaptor(shippingDto *domain.ShippingDto) error {
	systemsHandler := internalHandler.NewInternalServicesFacade()
	newShipping := systemsHandler.shippingService.BuildNewShipping(*shippingDto)
	if err := internalHandler.internal.CreateShippingAdaptor(*newShipping); err != nil {
		return err
	}
	pickUpDTO := domain.PickUpDto{
		ShippingID: newShipping.ID,
		CarrierID: newShipping.CarrierID,
		Status: string(domain.SCHEDULED),
		PickUpAt: time.Now(),
	}
	if err := internalHandler.NewPickUpAdaptor(pickUpDTO); err != nil {
		return err
	}
	return nil
}

func (internalHandler *InternalServicesHandler) GetShippingsAdaptor(ID uuid.UUID) ([]domain.Shipping, error) {
	shippings, err := internalHandler.internal.GetShippingsAdaptor(ID, "SCHEDULED")
	if err != nil {
		return []domain.Shipping{}, err
	}
	return shippings, nil
}
