package handler

import (
	"github.com/QUDUSKUNLE/microservices/diagnostic-service/core/services"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/diagnostic"
)

type DiagnosticHandler struct {
	DiagnosticService *services.DiagnosticService
	diagnostic.UnimplementedDiagnosticServiceServer
}

// NewDiagnosticHandler creates a new DiagnosticHandler with the given DiagnosticService.
// This function initializes the DiagnosticHandler with the provided DiagnosticService.
// It returns a pointer to the newly created DiagnosticHandler instance.
func NewDiagnosticHandler(diagnosticService *services.DiagnosticService) *DiagnosticHandler {
	return &DiagnosticHandler{
		DiagnosticService: diagnosticService,
	}
}
