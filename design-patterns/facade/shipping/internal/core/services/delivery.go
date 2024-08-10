package services

import (
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (internalHandler *InternalServicesHandler) NewDeliveryAdaptor(accountID uuid.UUID, productType domain.ProductType) error {
	return nil
}

func (internalHandler *InternalServicesHandler) NewDelivery(accountID uuid.UUID, pickUpAddress, deliveryAddress, productType string) error {
	systemsHandler := internalHandler.NewInternalServicesFacade()
	systemsHandler.notificationService.SendDeliveryNotification()
	return nil
}
