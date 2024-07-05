package schedule

import (
	"fmt"
)

type SchedulePickUp struct {}

func (pick *SchedulePickUp) SchedulePickUp(shippingID, pickUpAddress, deliveryAddress, date, tim string) {
	fmt.Printf("Product with shippingID %s has been pick up from pickUpAddress at %s\n and to be delivered %s on %s at time %s.\n", shippingID, pickUpAddress, deliveryAddress, date, tim)
}
