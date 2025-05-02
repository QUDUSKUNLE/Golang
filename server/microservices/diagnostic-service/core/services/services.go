package services

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/db"
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
	diagnostics, err := s.Repo.GetAllDiagnostics(ctx, db.GetAllDiagnosticsParams{Limit: 50, Offset: 0})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch diagnostics centers: %v", err)
	}
	userLat, userLon := req.GetLatitude(), req.GetLongitude()
	var responseDiagnostics []*diagnostic.Diagnostic
	for _, diag := range diagnostics {
		lat := diag.Latitude.Float64
		lon := diag.Longitude.Float64
		_, err := utils.Haversine(userLat, userLon, lat, lon)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid latitude or longitude: %v", err)
		}
		responseDiagnostics = append(responseDiagnostics, &diagnostic.Diagnostic{
			Id:        diag.ID,
			UserId:    diag.UserID,
			Name:      diag.Name.String,
			CreatedAt: diag.CreatedAt.Time.String(),
			UpdatedAt: diag.UpdatedAt.Time.String(),
		})
	}
	// Sort result by distance
	return &diagnostic.SearchNearestDiagnosticsResponse{
		Diagnostics: responseDiagnostics,
	}, nil
}

func (s *DiagnosticService) CancelDiagnostic(ctx context.Context, req *diagnostic.DeleteDiagnosticRequest) (*diagnostic.DeleteDiagnosticResponse, error) {
	// Delete the diagnostic from the database
	diag, err := s.Repo.CancelDiagnostic(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete diagnostic: %v", err)
	}
	return &diagnostic.DeleteDiagnosticResponse{
		Id:        diag.ID,
		UserId:    diag.UserID,
		CreatedAt: diag.CreatedAt.Time.String(),
		UpdatedAt: diag.UpdatedAt.Time.String(),
	}, nil
}

func (s *DiagnosticService) GetDiagnostic(ctx context.Context, req *diagnostic.GetDiagnosticRequest) (*diagnostic.GetDiagnosticResponse, error) {
	// Get the diagnostic from the database
	diag, err := s.Repo.GetDiagnostic(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch diagnostic: %v", err)
	}
	return &diagnostic.GetDiagnosticResponse{
		Id:        diag.ID,
		UserId:    diag.UserID,
		CreatedAt: diag.CreatedAt.Time.String(),
		UpdatedAt: diag.UpdatedAt.Time.String(),
	}, nil
}

func (s *DiagnosticService) ListDiagnostics(ctx context.Context, req *diagnostic.ListDiagnosticsRequest) (*diagnostic.ListDiagnosticsResponse, error) {
	// List diagnostics from the database
	diagnostics, err := s.Repo.ListDiagnostics(ctx, db.ListDiagnosticsParams{Limit: 50, Offset: 0})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch diagnostics: %v", err)
	}
	var responseDiagnostics []*diagnostic.Diagnostic
	for _, diag := range diagnostics {
		responseDiagnostics = append(responseDiagnostics, &diagnostic.Diagnostic{
			Id:        diag.ID,
			UserId:    diag.UserID,
			Name:      diag.Name.String,
			CreatedAt: diag.CreatedAt.Time.String(),
			UpdatedAt: diag.UpdatedAt.Time.String(),
		})
	}
	return &diagnostic.ListDiagnosticsResponse{
		Diagnostics: responseDiagnostics,
	}, nil
}
