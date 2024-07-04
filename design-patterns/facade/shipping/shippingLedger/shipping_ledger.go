package shipping

import (
	"fmt"
		product "github.com/QUDUSKUNLE/Golang/tutorial/design-patterns/facade/shipping/product"
)

type ShippingLedger struct {}

func (ledger *ShippingLedger) Ledger(accountID string, prod product.ProductType) (shippingID string, err error) {
	fmt.Printf("Make ledger entry for accountID %s with productType %s.\n", accountID, prod)
	return "1", nil
}
