package schedule

import (
	"fmt"
)

type ScheduleShipping struct {}

func (pick *ScheduleShipping) ScheduleShipping(shippingID, pickUpAddress, deliveryAddress, date, tim string) {
	fmt.Printf("Product with shippingID %s has been scheduled for shipping at %s\n and deliver to %s on %s at time %s.\n", shippingID, pickUpAddress, deliveryAddress, date, tim)
}
