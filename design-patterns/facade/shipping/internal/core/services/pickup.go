package services

import (
	"log"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (internalHandler *InternalServicesHandler) NewPickUpAdaptor(pick domain.PickUpDto) error {
	systemsHandler := internalHandler.NewInternalServicesFacade()
	pickUp := systemsHandler.pickUpService.BuildNewPickUp(pick)
	err := internalHandler.internal.InitiatePickUpAdaptor(*pickUp);
	if err != nil {
		return err
	}
	systemsHandler.notificationService.SendPickUpNotification()
	if err := systemsHandler.labelService.CreateShippingLabel(pickUp.ShippingID.String()); err != nil {
		log.Fatal(err)
	}
	// Send pick up notification
	systemsHandler.notificationService.SendPickUpNotification()
	return nil
}

func (internalHandler *InternalServicesHandler) UpDatePickUpAdaptor(pickUp domain.PickUpDto) error {
	systemsHandler := internalHandler.NewInternalServicesFacade()
	// build a new pick up
	pick := systemsHandler.pickUpService.BuildUpdatePickUp(pickUp)

	// Update pcik up ledger
	if err := internalHandler.internal.UpdatePickUpAdaptor(*pick); err != nil {
		return err
	}
	// Send pick up notification
	systemsHandler.notificationService.SendPickUpNotification()
	return nil
}
