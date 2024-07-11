package services

import (
	"fmt"
	"github.com/QUDUSKUNLE/shipping/internal/core/ledger"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/QUDUSKUNLE/shipping/internal/core/utils"
	"github.com/labstack/echo/v4"
)

type PickUpAdaptor struct {
	pickUpService *domain.PickUp
	pickUpRepositoryService *ledger.Ledger
	utilsService *utils.Utils
	notificationService *Notification
}

func NewPickUpAdaptor(pick domain.PickUpDTO) error {
	fmt.Println("Initiate a new pick up")
	adaptor := &PickUpAdaptor{
		pickUpService: &domain.PickUp{},
		pickUpRepositoryService: &ledger.Ledger{},
		notificationService: &Notification{},
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

func UpDatePickUpAdaptor(context echo.Context, pickUp domain.PickUp) error {
	fmt.Println("Update a parcel pick up")
	adaptor := &PickUpAdaptor{
		pickUpService: &domain.PickUp{},
		pickUpRepositoryService: &ledger.Ledger{},
		notificationService: &Notification{},
		utilsService: &utils.Utils{},
	}
	// Validate carrier
	carrier, err := adaptor.utilsService.ParseUserID(context)
	if err != nil {
		return err
	}

	if carrier.UserType != string(domain.RIDER) {
		return fmt.Errorf("unauthorized to perform this operation")
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
