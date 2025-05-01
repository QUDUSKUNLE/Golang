package handler

import (
	"github.com/QUDUSKUNLE/microservices/diagnostic-service/core/services"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/diagnostic"
)

type DiagnosticHandler struct {
	DiagnosticService *services.DiagnosticService
	diagnostic.UnimplementedDiagnosticServiceServer
}

func NewDiagnosticHandler(diagnosticService *services.DiagnosticService) *DiagnosticHandler {
	return &DiagnosticHandler{
		DiagnosticService: diagnosticService,
	}
}
