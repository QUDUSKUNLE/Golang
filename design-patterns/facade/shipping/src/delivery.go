package shipping

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/QUDUSKUNLE/shipping/src/notification"
	"github.com/QUDUSKUNLE/shipping/src/product"
	"github.com/QUDUSKUNLE/shipping/src/schedule"
	"github.com/QUDUSKUNLE/shipping/src/ledger"
)

type DeliveryAdaptor struct {
	user *model.User
	product *product.Product
	deliveryLedger *ledger.DeliveryLedger
	delivery *schedule.ScheduleDelivery
	notification *notification.Notification
}

func NewDeliveryAdaptor(accountID uuid.UUID, productType product.ProductType) *DeliveryAdaptor {
	fmt.Println("Initiate a new delivery")
	delivery :=  &DeliveryAdaptor{
		user: model.NewUser(accountID),
		product: product.NewProduct(productType),
		deliveryLedger: &ledger.DeliveryLedger{},
		delivery: &schedule.ScheduleDelivery{},
		notification: &notification.Notification{},
	}
	fmt.Println("New delivery initiated successfully.")
	return delivery
}

func (delivery *DeliveryAdaptor) NewDelivery(accountID uuid.UUID, pickUpAddress, deliveryAddress, productType string) error {
	fmt.Println("Start a new delivery.")
	if err := delivery.user.CheckUser(accountID); err != nil {
		return err
	}
	if err := delivery.product.CheckProduct(productType); err != nil {
		return err
	}
	deliveryID, err := delivery.deliveryLedger.Ledger(accountID, productType)
	if err != nil {
		return err;
	}
	delivery.delivery.ScheduleDelivery(deliveryID, pickUpAddress, deliveryAddress, "date", "time")
	delivery.notification.SendDeliveryNotification()
	fmt.Println("Product is delivered successfully.")
	return nil
}
