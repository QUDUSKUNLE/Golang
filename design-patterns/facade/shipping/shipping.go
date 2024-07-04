package shipping

import (
	account "github.com/QUDUSKUNLE/Golang/tutorial/design-patterns/facade/shipping/account"
	ledger "github.com/QUDUSKUNLE/Golang/tutorial/design-patterns/facade/shipping/shippingLedger"
	notification "github.com/QUDUSKUNLE/Golang/tutorial/design-patterns/facade/shipping/notification"
	schedulePickUp "github.com/QUDUSKUNLE/Golang/tutorial/design-patterns/facade/shipping/schedule"
	product "github.com/QUDUSKUNLE/Golang/tutorial/design-patterns/facade/shipping/product"

	"fmt"
)

type Shipping struct {
	account *account.User
	product *product.Product
	shippingLedger *ledger.ShippingLedger
	schedulePickUp *schedulePickUp.Schedule
	notification *notification.Notification
}


func NewShipping(accountID string, produc product.ProductType) *Shipping {
	fmt.Println("Start creating a new pickup shipping")
	ship := &Shipping{
		account: account.NewAccount(accountID),
		product: product.NewProduct(produc),
		shippingLedger: &ledger.ShippingLedger{},
		schedulePickUp: &schedulePickUp.Schedule{},
		notification: &notification.Notification{},
	}
	fmt.Println("Pickup shipping created successfully.")
	return ship
}

func (shipping *Shipping) ScheduleShipping(accountID, pickUpAddress string, prod product.ProductType) error {
	fmt.Println("Start to schedule a pickup shipping.")
	if err := shipping.account.CheckAccount(accountID); err != nil {
		return err
	}
	if err := shipping.product.CheckProduct(prod); err != nil {
		return err
	}
	shippingID, err := shipping.shippingLedger.Ledger(accountID, prod)
	if err != nil {
		return err;
	}
	shipping.schedulePickUp.SchedulePickUp(shippingID, pickUpAddress, "delivered", "date", "time")
	shipping.notification.SendShippingNotification()
	fmt.Println("Schedule pickup is successfull.")
	return nil
}
