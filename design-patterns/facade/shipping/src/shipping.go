package shipping

import (
	"github.com/QUDUSKUNLE/shipping/src/ledger"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/QUDUSKUNLE/shipping/src/notification"
	"github.com/QUDUSKUNLE/shipping/src/schedule"
	"github.com/QUDUSKUNLE/shipping/src/utils"
	"github.com/google/uuid"
	"time"
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

func (ship *ShippingAdaptor) NewShipping(ID uuid.UUID, dto model.ShippingDTO) error {
	// Build a new shipping
	newShipping, err := ship.shipping.BuildShipping(ID, dto)
	if err != nil {
		return err
	}
	// Log shipping request to shipping ledger
	err = ship.scheduleShippingLedger.ShippingLedger(newShipping)
	if err != nil {
		return err;
	}
	err = ship.schedulePickUp.SchedulePickUp(newShipping.ID, dto.PickUpAddress.State, dto.DeliveryAddress.State, time.Now().String(), time.Now().String())
	if err != nil {
		return err
	}
	ship.notification.SendShippingNotification()
	fmt.Println("Schedule pickup is successfull.")
	return nil
}
