package services

import (
	"context"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/diagnostic"
)

func (s *DiagnosticService) CreateDiagnostic(ctx context.Context, req *diagnostic.CreateDiagnosticRequest) (*diagnostic.CreateDiagnosticResponse, error) {
	// Save the diagnostic to the database
	diag, err := s.Repo.CreateDiagnostic(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &diagnostic.CreateDiagnosticResponse{
		Id: diag.ID,
		UserId: diag.UserID,
		CreatedAt: diag.CreatedAt.Time.String(),
		UpdatedAt: diag.UpdatedAt.Time.String(),
	}, nil
}
