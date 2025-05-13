package subscribe

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/QUDUSKUNLE/microservices/notification-service/emails"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
)

// Improved: handles event types, logs errors, and sends registration emails if needed
func SubscribeNotification(ctx context.Context, event []byte) error {
	var notificationEvent dto.NotificationEvent
	if err := json.Unmarshal(event, &notificationEvent); err != nil {
		log.Printf("failed to unmarshal event: %v", err)
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	log.Printf("Received notification event: %v", notificationEvent.EventType)

	switch notificationEvent.EventType {
	case "registration":
		// Example: send registration email if event type is registration
		data := notificationEvent.Data
		name := data["name"]
		email := data["email"]
		token := data["token"]
		if name == "" || email == "" || token == "" {
			log.Printf("missing registration data: name=%v, email=%v, token=%v", name, email, token)
			return fmt.Errorf("missing registration data")
		}
		if name == "" || email == "" || token == "" {
			log.Printf("missing registration data: name=%v, email=%v, token=%v", name, email, token)
			return fmt.Errorf("missing registration data")
		}

		// Load SMTP config from environment variables (or config management)
		cfg := emails.SMTPConfig{
			Host:     os.Getenv("SMTP_HOST"),
			Port:     587, // or parse from env
			Username: os.Getenv("SMTP_USER"),
			Password: os.Getenv("SMTP_PASS"),
			From:     os.Getenv("SMTP_FROM"),
		}
		ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := emails.SendRegistrationEmail(ctxTimeout, cfg, name, email, token); err != nil {
			log.Printf("failed to send registration email: %v", err)
			return err
		}
		log.Printf("Registration email sent to %s", email)
	default:
		log.Printf("Unhandled notification event type: %v", notificationEvent.EventType)
	}

	return nil
}
