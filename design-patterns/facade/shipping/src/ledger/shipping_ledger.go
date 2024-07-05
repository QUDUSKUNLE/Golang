package ledger

import (
	"fmt"
	"github.com/google/uuid"
)

type ScheduleShippingLedger struct {}

func (ledger *ScheduleShippingLedger) Ledger(accountID uuid.UUID, product string) (shippingID string, err error) {
	fmt.Printf("Make schedule shipping ledger entry for accountID %s with productType %s.\n", accountID, product)
	return "1", nil
}
