package subscribe

import (
	"context"
	"encoding/json"
	"fmt"

	// "log"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
)

func SubsribeNotification(ctx context.Context, event []byte) error {
	// Load configuration
	// Load environment variable
	// _, err := utils.LoadConfig()
	// if err != nil {
	// 	log.Fatalf("Error loading config: %v", err)
	// }
	// Process the event
	var notificationEvent dto.NotificationEvent
	if err := json.Unmarshal(event, &notificationEvent); err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	fmt.Printf("Process notification events: %v", notificationEvent.EventType)
	return nil
}
