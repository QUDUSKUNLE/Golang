package shipping

import (
	"fmt"
)

type Schedule struct {}

func (pick *Schedule) SchedulePickUp(shippingID, pickUpAddress, deliveryAddress, date, tim string) {
	fmt.Printf("Product with shippingID %s has been scheduled for pickUpAddress at %s\n and deliver to %s on %s at time %s.\n", shippingID, pickUpAddress, deliveryAddress, date, tim)
}


