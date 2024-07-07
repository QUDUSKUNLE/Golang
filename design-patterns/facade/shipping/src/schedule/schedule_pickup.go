package schedule

import (
		"github.com/google/uuid"
	"fmt"
)

type SchedulePickUp struct {}

func (pick *SchedulePickUp) SchedulePickUp(shippingID uuid.UUID, pickUpAddress, deliveryAddress, date, tim string) error {
	fmt.Printf("Product with shippingID %s has been scheduled for pick up from pickUpAddress at %s\n and to be delivered %s on %s at time %s.\n", shippingID, pickUpAddress, deliveryAddress, date, tim)
	return nil
}
