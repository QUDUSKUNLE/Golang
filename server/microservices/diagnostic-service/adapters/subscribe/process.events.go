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
)

func ProcessEvent(ctx context.Context, event []byte) error {
	// Process the event
	var user dto.UserCreatedEvent
	if err := json.Unmarshal(event, &user); err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	log.Printf("Processing DiagnosticCreatedEvent: UserID=%s", user.UserID)

	// Initialize the diagnostic service
	// repository.DiagnosticRepository
	diagnosticRepo := repository.NewDiagnosticRepository(db.DatabaseConnection())
	service := diagnosticService.NewDiagnosticService(*diagnosticRepo)
	diag, err := service.Repo.CreateDiagnostic(ctx, user.UserID)
	if err != nil {
		return fmt.Errorf("Error creating diagnostic centre: %w", err)
	}
	fmt.Printf("Diagnostic centre created successfully: %v", diag.ID)
	return nil
}
