package ledger

import (
	"fmt"

	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/google/uuid"
)

type ScheduleShippingLedger struct {}

func (ledger *ScheduleShippingLedger) ShippingLedger(userID uuid.UUID, dto model.ShippingDTO) (shippingID string, err error) {
	fmt.Println("Make schedule shipping ledger entry for accountID with productType")
	return "1", nil
}
