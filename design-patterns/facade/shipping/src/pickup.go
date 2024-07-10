package shipping

import (
	"fmt"

	"github.com/QUDUSKUNLE/shipping/src/ledger"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/QUDUSKUNLE/shipping/src/notification"
)

type PickUpAdaptor struct {
	userRepository *ledger.UserRepository
	pickUpService *model.PickUp
	pickUpRepositoryService *ledger.PickUpRepository
	notificationService *notification.Notification
}

func NewPickUpAdaptor(pick model.PickUpDTO) error {
	fmt.Println("Initiate a new pick up")
	adaptor := &PickUpAdaptor{
		pickUpService: &model.PickUp{},
		pickUpRepositoryService: &ledger.PickUpRepository{},
		notificationService: &notification.Notification{},
	}
	pickUp := adaptor.pickUpService.BuildNewPickUp(pick)
	err := adaptor.pickUpRepositoryService.NewLedger(*pickUp);
	if err != nil {
		fmt.Println(err)
		return err
	}
	adaptor.notificationService.SendPickUpNotification()
	fmt.Println("Parcel pickup initiated successfully.")
	return nil
}

func UpDatePickUpAdaptor(pickUp model.PickUp) error {
	fmt.Println("Update a parcel pick up")
	adaptor := &PickUpAdaptor{
		pickUpService: &model.PickUp{},
		userRepository: &ledger.UserRepository{},
		pickUpRepositoryService: &ledger.PickUpRepository{},
		notificationService: &notification.Notification{},
	}
	// Validate carrier
	_, err := adaptor.userRepository.QueryUserByID(pickUp.UserID)
	if err != nil {
		return err
	}
	// build a new pick up
	pick := adaptor.pickUpService.BuildUpdatePickUp(pickUp)

	// Update pcik up ledger
	err = adaptor.pickUpRepositoryService.UpdateLedger(*pick)
	if err != nil {
		return err
	}
	// Send pick up notification
	adaptor.notificationService.SendPickUpNotification()
	fmt.Println("Parcel pickup updated successfully.")
	return nil
}
