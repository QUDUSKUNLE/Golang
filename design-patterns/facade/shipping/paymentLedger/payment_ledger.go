package shipping

import "fmt"

type PaymentLedger struct {}

func (ledger *PaymentLedger) Payment(accountID, transactionType string, amount int) {
	fmt.Printf("Make ledger entry for accountID %s with transactionType %s for amount %d\n", accountID, transactionType, amount)
}
