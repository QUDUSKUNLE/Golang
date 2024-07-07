package shipping

import (
	"fmt"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/QUDUSKUNLE/shipping/src/notification"
)

type PickUpAdaptor struct {
	pickUpService *model.PickUp
	notificationService *notification.Notification
}

func NewPickUpAdaptor(pick model.PickUpDTO) error {
	fmt.Println("Initiate a new product pick up")
	adaptor := &PickUpAdaptor{
		pickUpService: &model.PickUp{},
		notificationService: &notification.Notification{},
	}
	pickUp := adaptor.pickUpService.BuildNewPickUp(pick)
	_, err := adaptor.pickUpService.NewLedger(*pickUp);
	if err != nil {
		return err
	}
	adaptor.notificationService.SendPickUpNotification()
	return nil
}

func UpDatePickUpAdaptor(pickUp model.PickUpDTO) error {
	fmt.Println("Update a parcel pick up")
	adaptor := &PickUpAdaptor{
		pickUpService: &model.PickUp{},
		notificationService: &notification.Notification{},
	}
	pick := adaptor.pickUpService.BuildUpdatePickUp(pickUp)
	_, err := adaptor.pickUpService.UpdateLedger(*pick)
	if err != nil {
		return err
	}
	adaptor.notificationService.SendPickUpNotification()
	fmt.Println("Parcel pickup updated successfully.")
	return nil
}
