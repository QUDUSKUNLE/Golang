package services

import (
	"fmt"
	"log"

	"github.com/wneessen/go-mail"
)

type NotificationService struct {}

func (notification *NotificationService) emailNotification() (err error){
	msg := mail.NewMsg()
	if err := msg.From("okay@tester.com"); err != nil {
		log.Fatalf("failed to set FROM address: %s", err)
	}
	if err := msg.To("quduskunle@fgmail.com"); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}
	msg.Subject("This is my first test mail with go-mail")
	msg.SetBodyString(mail.TypeTextPlain, "This will be content of the mail.")
	fmt.Println("Sening.... email")
	return
}

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
	if err := notification.emailNotification(); err != nil {
		log.Fatalf("%s", err)
	}
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
	if err := notification.emailNotification(); err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println("Sending password reset notification...")
}
