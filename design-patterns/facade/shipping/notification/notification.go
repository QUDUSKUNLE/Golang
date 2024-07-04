package shipping

import "fmt"

type Notification struct {}

func (notification *Notification) SendShippingNotification() {
	fmt.Println("Sending shipping notification.")
}

func (notification *Notification) SendPickUpNotification() {
	fmt.Println("Sending pickup notification.")
}

func (notification *Notification) SendReturnNotification() {
	fmt.Println("Sending return notification.")
}

func (notification *Notification) SendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (notification *Notification) SendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}
