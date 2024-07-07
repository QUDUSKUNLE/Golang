package shipping

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/QUDUSKUNLE/shipping/src/notification"
	"github.com/QUDUSKUNLE/shipping/src/schedule"
	"github.com/QUDUSKUNLE/shipping/src/ledger"
)

type DeliveryAdaptor struct {
	user *model.User
	shipping *model.Shipping
	deliveryLedger *ledger.DeliveryLedger
	delivery *schedule.ScheduleDelivery
	notification *notification.Notification
}

func NewDeliveryAdaptor(accountID uuid.UUID, productType model.ProductType) *DeliveryAdaptor {
	fmt.Println("Initiate a new delivery")
	delivery :=  &DeliveryAdaptor{
		user: &model.User{},
		shipping: &model.Shipping{},
		deliveryLedger: &ledger.DeliveryLedger{},
		delivery: &schedule.ScheduleDelivery{},
		notification: &notification.Notification{},
	}
	fmt.Println("New delivery initiated successfully.")
	return delivery
}

func (delivery *DeliveryAdaptor) NewDelivery(accountID uuid.UUID, pickUpAddress, deliveryAddress, productType string) error {
	fmt.Println("Start a new delivery.")
	deliveryID, err := delivery.deliveryLedger.Ledger(accountID, productType)
	if err != nil {
		return err;
	}
	delivery.delivery.ScheduleDelivery(deliveryID, pickUpAddress, deliveryAddress, "date", "time")
	delivery.notification.SendDeliveryNotification()
	fmt.Println("Product is delivered successfully.")
	return nil
}
