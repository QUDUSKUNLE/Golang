package shipping

import (
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/QUDUSKUNLE/shipping/src/ledger"
	"github.com/QUDUSKUNLE/shipping/src/notification"
	"github.com/QUDUSKUNLE/shipping/src/schedule"
	"github.com/QUDUSKUNLE/shipping/src/product"

	"fmt"
)

type ShippingAdaptor struct {
	user *model.User
	product *product.Product
	scheduleshippingLedger *ledger.ScheduleShippingLedger
	scheduleShipping *schedule.ScheduleShipping
	notification *notification.Notification
}


func NewShippingAdaptor(accountID uuid.UUID, productType product.ProductType) *ShippingAdaptor {
	fmt.Println("Start creating a new pickup shipping")
	ship := &ShippingAdaptor{
		user: model.NewUser(accountID),
		product: product.NewProduct(productType),
		scheduleshippingLedger: &ledger.ScheduleShippingLedger{},
		scheduleShipping: &schedule.ScheduleShipping{},
		notification: &notification.Notification{},
	}
	fmt.Println("Shipping created successfully.")
	fmt.Print("")
	return ship
}

func (shipping *ShippingAdaptor) NewShipping(accountID uuid.UUID, pickUpAddress, deliveryAddress, productType string) error {
	fmt.Println("Start to schedule a pickup shipping.")
	if err := shipping.user.CheckUser(accountID); err != nil {
		return err
	}
	if err := shipping.product.CheckProduct(productType); err != nil {
		return err
	}
	shippingID, err := shipping.scheduleshippingLedger.Ledger(accountID, productType)
	if err != nil {
		return err;
	}
	shipping.scheduleShipping.ScheduleShipping(shippingID, pickUpAddress, deliveryAddress, "date", "time")
	shipping.notification.SendShippingNotification()
	fmt.Println("Schedule pickup is successfull.")
	return nil
}
