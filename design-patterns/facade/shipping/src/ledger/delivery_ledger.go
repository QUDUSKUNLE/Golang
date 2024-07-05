package ledger

import (
	"fmt"
)

type DeliveryLedger struct {}

func (ledger *DeliveryLedger) Ledger(accountID, product string) (shippingID string, err error) {
	fmt.Printf("Make a Delivery ledger entry for accountID %s with productType %s.\n", accountID, product)
	return "1", nil
}
