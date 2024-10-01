package facade

import "fmt"

type Notification struct {}

func (notification *Notification) SendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (notification *Notification) SendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}
