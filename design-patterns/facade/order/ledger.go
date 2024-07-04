package facade

import "fmt"

type Ledger struct {}

func (ledger *Ledger) MakeEntry(accountID, transactionType string, amount int) {
	fmt.Printf("Make ledger entry for accountID %s with transactionType %s for amount %d\n", accountID, transactionType, amount)
}
