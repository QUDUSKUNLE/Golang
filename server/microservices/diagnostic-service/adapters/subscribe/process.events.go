package subscribe

import (
	"context"
	"encoding/json"
	"fmt"

	"log"

	"github.com/QUDUSKUNLE/microservices/diagnostic-service/adapters/repository"
	diagnosticService "github.com/QUDUSKUNLE/microservices/diagnostic-service/core/services"
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
	var user dto.UserCreatedEvent
	if err := json.Unmarshal(event, &user); err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	log.Printf("Processing DiagnosticCreatedEvent: UserID=%s", user.UserID)

	// Initialize the diagnostic service
	// repository.DiagnosticRepository
	diagnosticRepo := repository.NewDiagnosticRepository(db.DatabaseConnection(cfg.DB_URL))
	service := diagnosticService.NewDiagnosticService(*diagnosticRepo)
	diag, err := service.Repo.CreateDiagnostic(ctx, user.UserID)
	if err != nil {
		return fmt.Errorf("Error creating diagnostic centre: %w", err)
	}
	fmt.Printf("Diagnostic centre created successfully: %v", diag.ID)
	return nil
}
