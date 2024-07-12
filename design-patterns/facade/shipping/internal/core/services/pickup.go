package services

import (
	"fmt"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

type PickUpAdaptor struct {
	pickUpService *domain.PickUp
	notificationService *Notification
}

func (httpHandler *ServicesHandler) NewPickUpAdaptor(pick domain.PickUpDTO) error {
	fmt.Println("Initiate a new pick up")
	adaptor := &PickUpAdaptor{
		pickUpService: &domain.PickUp{},
		notificationService: &Notification{},
	}
	pickUp := adaptor.pickUpService.BuildNewPickUp(pick)
	err := httpHandler.Internal.InitiatePickUpAdaptor(*pickUp);
	if err != nil {
		fmt.Println(err)
		return err
	}
	adaptor.notificationService.SendPickUpNotification()
	fmt.Println("Parcel pickup initiated successfully.")
	return nil
}

func (httpHandler *ServicesHandler) UpDatePickUpAdaptor(pickUp domain.PickUp) error {
	fmt.Println("Update a parcel pick up")
	adaptor := &PickUpAdaptor{
		pickUpService: &domain.PickUp{},
		notificationService: &Notification{},
	}
	// build a new pick up
	pick := adaptor.pickUpService.BuildUpdatePickUp(pickUp)

	// Update pcik up ledger
	if err := httpHandler.Internal.UpdatePickUpAdaptor(*pick); err != nil {
		return err
	}
	// Send pick up notification
	adaptor.notificationService.SendPickUpNotification()
	fmt.Println("Parcel pickup updated successfully.")
	return nil
}
