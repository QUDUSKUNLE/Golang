package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

type DeliveryAdaptor struct {
	user *domain.User
	shipping *domain.Shipping
	notification *Notification
}

func (httpHandler *ServicesHandler) NewDeliveryAdaptor(accountID uuid.UUID, productType domain.ProductType) *DeliveryAdaptor {
	fmt.Println("Initiate a new delivery")
	delivery :=  &DeliveryAdaptor{
		user: &domain.User{},
		shipping: &domain.Shipping{},
		notification: &Notification{},
	}
	fmt.Println("New delivery initiated successfully.")
	return delivery
}

func (delivery *DeliveryAdaptor) NewDelivery(accountID uuid.UUID, pickUpAddress, deliveryAddress, productType string) error {
	fmt.Println("Start a new delivery.")
	delivery.notification.SendDeliveryNotification()
	fmt.Println("Product is delivered successfully.")
	return nil
}
