package services

import "fmt"

type NotificationService struct {}

func (notification *NotificationService) SendShippingNotification() {
	fmt.Println("Sending shipping notification.")
}

func (notification *NotificationService) SendPickUpNotification() {
	fmt.Println("Sending pickup notification.")
}

func (notification *NotificationService) SendReturnNotification() {
	fmt.Println("Sending return notification.")
}

func (notification *NotificationService) SendAddressNotification() {
	fmt.Println("Sending address notification.")
}

func (notification *NotificationService) SendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (notification *NotificationService) SendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

func (notification *NotificationService) SendDeliveryNotification() {
	fmt.Println("Sending delivery notification")
}

func (notification *NotificationService) SendRegistrationNotification() {
	fmt.Println("Sending registration notification")
}

func (notification *NotificationService) SendParcelNotification() {
	fmt.Println("Sending parcel notification")
}

func (notification *NotificationService) SendPackagingNotification() {
	fmt.Println("Sending packaging notification")
}

func (notification *NotificationService) SendResetPasswordNotification() {
	fmt.Println("Sending password reset notification")
}
