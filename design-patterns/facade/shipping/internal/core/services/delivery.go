package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (httpHandler *ServicesHandler) NewDeliveryAdaptor(accountID uuid.UUID, productType domain.ProductType) error {
	fmt.Println("Initiate a new delivery")
	fmt.Println("New delivery initiated successfully.")
	return nil
}

func (httpHandler *ServicesHandler) NewDelivery(accountID uuid.UUID, pickUpAddress, deliveryAddress, productType string) error {
	fmt.Println("Start a new delivery.")
	systemsHandler := httpHandler.NewServicesFacade()
	systemsHandler.notificationService.SendDeliveryNotification()
	fmt.Println("Product is delivered successfully.")
	return nil
}
