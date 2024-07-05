package schedule

import (
	"fmt"
)

type ScheduleDelivery struct {}

func (pick *ScheduleDelivery) ScheduleDelivery(shippingID, pickUpAddress, deliveryAddress, date, tim string) {
	fmt.Printf("Product with shippingID %s has been delivered from pickUpAddress at %s\n and delivered to %s on %s at time %s.\n", shippingID, pickUpAddress, deliveryAddress, date, tim)
}
