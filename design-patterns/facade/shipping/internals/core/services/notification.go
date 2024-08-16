package services

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

func (notification *Notification) SendAddressNotification() {
	fmt.Println("Sending address notification.")
}

func (notification *Notification) SendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (notification *Notification) SendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

func (notification *Notification) SendDeliveryNotification() {
	fmt.Println("Sending delivery notification")
}

func (notification *Notification) SendRegistrationNotification() {
	fmt.Println("Sending registration notification")
}

func (notification *Notification) SendParcelNotification() {
	fmt.Println("Sending parcel notification")
}

func (notification *Notification) SendPackagingNotification() {
	fmt.Println("Sending packaging notification")
}

func (notification *Notification) SendResetPasswordNotification() {
	fmt.Println("Sending password reset notification")
}
