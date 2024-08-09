package services

import (
	"fmt"
	"log"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (httpHandler *InternalServicesHandler) NewPickUpAdaptor(pick domain.PickUpDTO) error {
	fmt.Println("Initiate a new pick up")
	systemsHandler := httpHandler.NewInternalServicesFacade()
	pickUp := systemsHandler.pickUpService.BuildNewPickUp(pick)
	err := httpHandler.internal.InitiatePickUpAdaptor(*pickUp);
	if err != nil {
		return err
	}
	systemsHandler.notificationService.SendPickUpNotification()
	if err := systemsHandler.labelService.CreateShippingLabel(pickUp.ShippingID.String()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Parcel pickup initiated successfully.")
	return nil
}

func (httpHandler *InternalServicesHandler) UpDatePickUpAdaptor(pickUp domain.PickUp) error {
	fmt.Println("Update a parcel pick up")
	systemsHandler := httpHandler.NewInternalServicesFacade()
	// build a new pick up
	pick := systemsHandler.pickUpService.BuildUpdatePickUp(pickUp)

	// Update pcik up ledger
	if err := httpHandler.internal.UpdatePickUpAdaptor(*pick); err != nil {
		return err
	}
	// Send pick up notification
	systemsHandler.notificationService.SendPickUpNotification()
	fmt.Println("Parcel pickup updated successfully.")
	return nil
}
