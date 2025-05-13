package domain

const (
	// EmailNotificationType represents the type of email notification
	RegistrationEmail= "RegistrationEmail"
	ResetPasswordEmail = "ResetPasswordEmail"
)

type Notification interface {
	EmailNotification(userID, message string) error
	PushNotification(userID, message string) error
	SMSNotification(userID, message string) error
}

type NotificationService struct {
	// Add fields if necessary
}

// SendNotification sends a notification to a user
func (ns *NotificationService) EmailNotification(userID, message string) error {
	return nil
}

// GetNotifications retrieves notifications for a user
func (ns *NotificationService) PushNotification(userID string) error {
	// Implementation here
	return nil
}

// MarkAsRead marks a notification as read
func (ns *NotificationService) SMSNotification(userID, message string) error {
	// Implementation here
	return nil
}
