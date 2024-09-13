package services

import (
	"log"
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
)

func (internalHandler *InternalServicesHandler) NewPickUpAdaptor(pick domain.PickUpDto) error {
	systemsHandler := NewInternalServicesFacade()
	pickUp := systemsHandler.pickUpService.BuildNewPickUp(pick)
	err := internalHandler.internal.InitiatePickUpAdaptor(pickUp);
	if err != nil {
		return err
	}
	systemsHandler.notificationService.SendPickUpNotification()
	for _, pick := range pickUp {
		if err := systemsHandler.labelService.CreateShippingLabel(pick.ShippingID.String()); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func (internalHandler *InternalServicesHandler) UpDatePickUpAdaptor(pickUp domain.PickUpDto) error {
	systemsHandler := NewInternalServicesFacade()
	// build a new pick up
	pick := systemsHandler.pickUpService.BuildUpdatePickUp(pickUp)

	// Update pcik up ledger
	for _, p := range pick {
		if err := internalHandler.internal.UpdatePickUpAdaptor(*p); err != nil {
			log.Fatal(err)
		}

	}
	// Send pick up notification
	systemsHandler.notificationService.SendPickUpNotification()
	return nil
}

func (internalHandler *InternalServicesHandler) GetPickUpAdaptor(pickUpID, userID uuid.UUID) (*domain.PickUp, error) {
	// build a new pick up
	pickUp, err := internalHandler.internal.GetPickUp(pickUpID, userID)
	if err != nil {
		return &domain.PickUp{}, err
	}
	return &pickUp, nil
}
