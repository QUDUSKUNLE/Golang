package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (httpHandler *InternalServicesHandler) NewDeliveryAdaptor(accountID uuid.UUID, productType domain.ProductType) error {
	fmt.Println("Initiate a new delivery")
	fmt.Println("New delivery initiated successfully.")
	return nil
}

func (httpHandler *InternalServicesHandler) NewDelivery(accountID uuid.UUID, pickUpAddress, deliveryAddress, productType string) error {
	fmt.Println("Start a new delivery.")
	systemsHandler := httpHandler.NewInternalServicesFacade()
	systemsHandler.notificationService.SendDeliveryNotification()
	fmt.Println("Product is delivered successfully.")
	return nil
}
