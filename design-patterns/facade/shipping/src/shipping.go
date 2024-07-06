package shipping

import (
	"github.com/QUDUSKUNLE/shipping/src/ledger"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/QUDUSKUNLE/shipping/src/notification"
	"github.com/QUDUSKUNLE/shipping/src/schedule"
	"github.com/QUDUSKUNLE/shipping/src/utils"
	"github.com/google/uuid"

	"fmt"
)

type ShippingAdaptor struct {
	user *model.User
	shipping *model.Shipping
	scheduleShippingLedger *ledger.ScheduleShippingLedger
	schedulePickUp *schedule.SchedulePickUp
	notification *notification.Notification
	Utils *utils.Utils
}

func NewShippingAdaptor() *ShippingAdaptor {
	fmt.Println("Start creating a new pickup shipping")
	ship := &ShippingAdaptor{
		user: &model.User{},
		shipping: &model.Shipping{},
		scheduleShippingLedger: &ledger.ScheduleShippingLedger{},
		schedulePickUp: &schedule.SchedulePickUp{},
		notification: &notification.Notification{},
		Utils: &utils.Utils{},
	}
	fmt.Println("Shipping created successfully.")
	return ship
}

func (shipp *ShippingAdaptor) NewShipping(ID uuid.UUID, dto model.ShippingDTO) error {
	fmt.Println("Start to schedule a new shipping.")
	shippingID, err := shipp.scheduleShippingLedger.ShippingLedger(ID, dto)
	if err != nil {
		return err;
	}
	fmt.Println(shippingID)
	// shipp.schedulePickUp.SchedulePickUp(shippingID, pickUpAddress.State, deliveryAddress.State, "date", "time")
	shipp.notification.SendShippingNotification()
	fmt.Println("Schedule pickup is successfull.")
	return nil
}
