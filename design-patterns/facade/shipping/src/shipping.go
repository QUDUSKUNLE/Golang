package shipping

import (
	"github.com/QUDUSKUNLE/shipping/src/ledger"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/QUDUSKUNLE/shipping/src/notification"
	"github.com/QUDUSKUNLE/shipping/src/schedule"
	"github.com/QUDUSKUNLE/shipping/src/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	// "time"
	"fmt"
)

type ShippingAdaptor struct {
	utils *utils.Utils
	shipping *model.Shipping
	shippingRepository *ledger.ShippingRepository
	pickUp *model.PickUp
	user *model.User
	schedulePickUp *schedule.SchedulePickUp
	notification *notification.Notification
}

func NewShippingAdaptor(cont echo.Context, shippingDto *model.ShippingDTO) error {
	fmt.Println("Start a new pickup shipping")
	adaptor := &ShippingAdaptor{
		user: &model.User{},
		shipping: &model.Shipping{},
		shippingRepository: &ledger.ShippingRepository{},
		pickUp: &model.PickUp{},
		schedulePickUp: &schedule.SchedulePickUp{},
		notification: &notification.Notification{},
		utils: &utils.Utils{},
	}
	userID, err := uuid.Parse(adaptor.utils.ObtainUser(cont))
	if err != nil {
		return err
	}
	newShipping := adaptor.shipping.BuildNewShipping(userID, *shippingDto)
	err = adaptor.shippingRepository.ShippingLedger(*newShipping)
	if err != nil {
		return err
	}
	fmt.Println("Shipping created successfully.")
	return nil
}

func (ship *ShippingAdaptor) NewShipping(ID uuid.UUID, dto model.ShippingDTO) error {
	// Build a new shipping
	newShipping := ship.shipping.BuildNewShipping(ID, dto)
	// Log shipping request to shipping ledger
	err := ship.shippingRepository.ShippingLedger(*newShipping)
	if err != nil {
		return err;
	}
	// Alert Pick up service for scheduling
	// err = ship.schedulePickUp.SchedulePickUp(newShipping.ID, dto.PickUpAddress.State, dto.DeliveryAddress.State, time.Now().String(), time.Now().String())
	// if err != nil {
	// 	return err
	// }
	// Alert notification service
	ship.notification.SendShippingNotification()
	fmt.Println("Schedule pickup is successfull.")
	return nil
}
