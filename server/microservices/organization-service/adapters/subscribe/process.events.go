package subscribe

import (
	"context"
	"encoding/json"
	"fmt"

	"log"

	organizationService "github.com/QUDUSKUNLE/microservices/organization-service/adapters/organizationcase"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
)

func ProcessEvent(ctx context.Context, event []byte) error {
	// Load configuration
	// Load environment variable
	cfg, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	// Process the event
	// For example, you can unmarshal the event and perform some action based on its content
	var user dto.UserCreatedEvent
	if err := json.Unmarshal(event, &user); err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	log.Printf("Processing OrganizationCreatedEvent: UserID=%s", user.UserID)

    // Initialize the organization service
    organizationService := organizationService.InitOrganizationServer(db.DatabaseConnection(cfg.DB_URL))
    organization, err := organizationService.CreateOrganization(ctx, dto.OrganizationDto{UserID: user.UserID})
    if err != nil {
        return fmt.Errorf("error creating organization: %w", err)
    }

    log.Printf("Organization created successfully: %v", organization.ID)
    return nil
}
