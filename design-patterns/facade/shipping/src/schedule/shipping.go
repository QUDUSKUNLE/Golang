package schedule

import (
	"fmt"
	"github.com/QUDUSKUNLE/shipping/src/model"
)

type ScheduleShipping struct {}

func (pick *ScheduleShipping) ScheduleShipping(shippingID string, pickUpAddress, deliveryAddress model.Address, date, tim string) {
	fmt.Printf("Product with shippingID %s has been scheduled for shipping at %s\n and deliver to %s on %s at time %s.\n", shippingID, pickUpAddress.StreetName, deliveryAddress.StreetName, date, tim)
}
