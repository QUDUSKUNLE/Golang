package handler

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/protogen/diagnostic"
)

func (h *DiagnosticHandler) CreateDiagnostic(ctx context.Context, req *diagnostic.CreateDiagnosticRequest) (*diagnostic.CreateDiagnosticResponse, error) {
	return h.DiagnosticService.CreateDiagnostic(ctx, req)
}

func (h *DiagnosticHandler) GetDiagnostic(ctx context.Context, req *diagnostic.GetDiagnosticRequest) (*diagnostic.GetDiagnosticResponse, error) {
	return h.DiagnosticService.GetDiagnostic(ctx, req)
}

func (h *DiagnosticHandler) ListDiagnostics(ctx context.Context, req *diagnostic.ListDiagnosticsRequest) (*diagnostic.ListDiagnosticsResponse, error) {
	return h.DiagnosticService.ListDiagnostics(ctx, req)
}

func (h *DiagnosticHandler) UpdateDiagnostic(ctx context.Context, req *diagnostic.UpdateDiagnosticRequest) (*diagnostic.UpdateDiagnosticResponse, error) {
	return h.DiagnosticService.UpdateDiagnostic(ctx, req)
}

func (h *DiagnosticHandler) CancelDiagnostic(ctx context.Context, req *diagnostic.DeleteDiagnosticRequest) (*diagnostic.DeleteDiagnosticResponse, error) {
	return h.DiagnosticService.CancelDiagnostic(ctx, req)
}



func (h *DiagnosticHandler) SearchNearestDiagnosticCenter(ctx context.Context, req *diagnostic.SearchNearestDiagnosticsRequest) (*diagnostic.SearchNearestDiagnosticsResponse, error) {
	return h.DiagnosticService.SearchNearestDiagnosticCenter(ctx, req)
}
