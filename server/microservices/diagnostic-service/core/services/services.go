package services

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/protogen/diagnostic"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *DiagnosticService) CreateDiagnostic(ctx context.Context, req *diagnostic.CreateDiagnosticRequest) (*diagnostic.CreateDiagnosticResponse, error) {
	// Save the diagnostic to the database
	diag, err := s.Repo.CreateDiagnostic(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &diagnostic.CreateDiagnosticResponse{
		Id:        diag.ID,
		UserId:    diag.UserID,
		CreatedAt: diag.CreatedAt.Time.String(),
		UpdatedAt: diag.UpdatedAt.Time.String(),
	}, nil
}

func (s *DiagnosticService) SearchNearestDiagnosticCenter(ctx context.Context, req *diagnostic.SearchNearestDiagnosticsRequest) (*diagnostic.SearchNearestDiagnosticsResponse, error) {
	// Search for the nearest diagnostic center
	diagnostics, err := s.Repo.GetAllDiagnostics(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch diagnostics centers: %v", err)
	}
	userLat, userLon := req.GetLatitude(), req.GetLongitude()
	var responseDiagnostics []*diagnostic.Diagnostic
	for _, diag := range diagnostics {
		lat := diag.Latitude.Float64
		lon := diag.Longitude.Float64
		_ = utils.Haversine(userLat, userLon, lat, lon)
		responseDiagnostics = append(responseDiagnostics, &diagnostic.Diagnostic{
			Id:        diag.ID,
			UserId:    diag.UserID,
			CreatedAt: diag.CreatedAt.Time.String(),
			UpdatedAt: diag.UpdatedAt.Time.String(),
		})
	}
	// Sort result by distance
	return &diagnostic.SearchNearestDiagnosticsResponse{
		Diagnostics: responseDiagnostics,
	}, nil
}
