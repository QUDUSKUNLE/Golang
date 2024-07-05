package shipping

import (
	"fmt"
	"github.com/QUDUSKUNLE/shipping/src/account"
	"github.com/QUDUSKUNLE/shipping/src/ledger"
	"github.com/QUDUSKUNLE/shipping/src/notification"
	"github.com/QUDUSKUNLE/shipping/src/product"
	"github.com/QUDUSKUNLE/shipping/src/schedule"
)

type PickUp struct {
	user *account.User
	product *product.Product
	pickUpLedger *ledger.PickUpLedger
	schedulePickUp *schedule.SchedulePickUp
	notification *notification.Notification
}

func NewPickUp(accountID string, productType product.ProductType) *PickUp {
	fmt.Println("Initiate a new product pick up")
	pickup := &PickUp{
		user: account.NewUser(accountID),
		product: product.NewProduct(productType),
		pickUpLedger: &ledger.PickUpLedger{},
		schedulePickUp: &schedule.SchedulePickUp{},
		notification: &notification.Notification{},
	}
	fmt.Println("Product picked up initiated successfully.")
	return pickup
}


func (pickUp *PickUp) NewSchedulePickUp(accountID, pickUpAddress, deliveryAddress, productType string) error {
	fmt.Println("Start a new pickup.")
	if err := pickUp.user.CheckUser(accountID); err != nil {
		return err
	}
	if err := pickUp.product.CheckProduct(productType); err != nil {
		return err
	}
	shippingID, err := pickUp.pickUpLedger.Ledger(accountID, productType)
	if err != nil {
		return err;
	}
	pickUp.schedulePickUp.SchedulePickUp(shippingID, pickUpAddress, deliveryAddress, "date", "time")
	pickUp.notification.SendPickUpNotification()
	fmt.Println("Product is picked up successfully.")
	return nil
}
