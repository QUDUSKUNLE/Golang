package handler

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/protogen/diagnostic"
)

func (h *DiagnosticHandler) CreateDiagnostic(ctx context.Context, req *diagnostic.CreateDiagnosticRequest) (*diagnostic.CreateDiagnosticResponse, error) {
	return h.DiagnosticService.CreateDiagnostic(ctx, req)
}

func (h *DiagnosticHandler) SearchNearestDiagnosticCenter(ctx context.Context, req *diagnostic.SearchNearestDiagnosticsRequest) (*diagnostic.SearchNearestDiagnosticsResponse, error) {
	return h.DiagnosticService.SearchNearestDiagnosticCenter(ctx, req)
}
