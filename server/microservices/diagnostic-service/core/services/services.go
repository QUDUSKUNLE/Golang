package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/QUDUSKUNLE/microservices/shared/constants"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/middleware"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/diagnostic"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

func (s *DiagnosticService) CreateDiagnostic(ctx context.Context, req *diagnostic.CreateDiagnosticRequest) (*diagnostic.CreateDiagnosticResponse, error) {
	// Validate the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumDIAGNOSTICCENTRE))
	if err != nil {
		utils.LogError("Error validating diagnostic centre: ", err)
		return nil, status.Errorf(codes.PermissionDenied, "Unauthorized: %v", err)
	}
	// Save the diagnostic to the database
	diag, err := s.Repo.CreateDiagnostic(ctx, db.CreateDiagnosticParams{
		UserID: user.UserID, DiagnosticCentreName: req.GetDiagnosticCentreName()})
	if err != nil {
		return nil, err
	}
	return &diagnostic.CreateDiagnosticResponse{
		DiagnosticId: diag.ID,
		UserId:       diag.UserID,
		CreatedAt:    diag.CreatedAt.Time.String(),
		UpdatedAt:    diag.UpdatedAt.Time.String(),
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
			DiagnosticId:         diag.ID,
			UserId:               diag.UserID,
			DiagnosticCentreName: diag.DiagnosticCentreName,
			CreatedAt:            diag.CreatedAt.Time.String(),
			UpdatedAt:            diag.UpdatedAt.Time.String(),
		})
	}
	// Sort result by distance
	return &diagnostic.SearchNearestDiagnosticsResponse{
		Diagnostics: responseDiagnostics,
	}, nil
}

func (s *DiagnosticService) CancelDiagnostic(ctx context.Context, req *diagnostic.DeleteDiagnosticRequest) (*diagnostic.DeleteDiagnosticResponse, error) {
	// Validate the request
	_, err := middleware.ValidateUser(ctx, string(db.UserEnumDIAGNOSTICCENTRE))
	if err != nil {
		utils.LogError("Error validating diagnostic centre: ", err)
		return nil, status.Errorf(codes.PermissionDenied, "Unauthorized: %v", err)
	}

	// Delete the diagnostic from the database
	diag, err := s.Repo.CancelDiagnostic(ctx, req.GetDiagnosticId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete diagnostic: %v", err)
	}
	return &diagnostic.DeleteDiagnosticResponse{
		DiagnosticId: diag.ID,
		UserId:       diag.UserID,
		CreatedAt:    diag.CreatedAt.Time.String(),
		UpdatedAt:    diag.UpdatedAt.Time.String(),
	}, nil
}

func (s *DiagnosticService) GetDiagnostic(ctx context.Context, req *diagnostic.GetDiagnosticRequest) (*diagnostic.GetDiagnosticResponse, error) {
	// Validate the request
	_, err := middleware.ValidateUser(ctx, string(db.UserEnumDIAGNOSTICCENTRE))
	if err != nil {
		utils.LogError("Error validating diagnostic centre: ", err)
		return nil, status.Errorf(codes.PermissionDenied, "Unauthorized: %v", err)
	}
	// Get the diagnostic from the database
	diag, err := s.Repo.GetDiagnostic(ctx, req.GetDiagnosticId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch diagnostic: %v", err)
	}
	return &diagnostic.GetDiagnosticResponse{
		DiagnosticId:         diag.ID,
		UserId:               diag.UserID,
		DiagnosticCentreName: diag.DiagnosticCentreName,
		Latitude:             diag.Latitude.Float64,
		Longitude:            diag.Longitude.Float64,
		Address: func() *structpb.Struct {
			var addressStruct structpb.Struct
			if err := addressStruct.UnmarshalJSON(diag.Address); err != nil {
				utils.LogError("Failed to unmarshal address: ", err)
				return nil
			}
			return &addressStruct
		}(),
		Contact: func() *structpb.Struct {
			var contactStruct structpb.Struct
			if err := contactStruct.UnmarshalJSON(diag.Contact); err != nil {
				utils.LogError("Failed to unmarshal contact: ", err)
				return nil
			}
			return &contactStruct
		}(),
		CreatedAt: diag.CreatedAt.Time.String(),
		UpdatedAt: diag.UpdatedAt.Time.String(),
	}, nil
}

func (s *DiagnosticService) ListDiagnostics(ctx context.Context, req *diagnostic.ListDiagnosticsRequest) (*diagnostic.ListDiagnosticsResponse, error) {
	// Validate the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumDIAGNOSTICCENTRE))
	if err != nil {
		utils.LogError("Error validating diagnostic centre: ", err)
		return nil, status.Errorf(codes.PermissionDenied, "Unauthorized: %v", err)
	}
	// List diagnostics from the database
	diagnostics, err := s.Repo.ListDiagnostics(
		ctx,
		db.ListDiagnosticsParams{
			Limit: func() int32 {
				if req.GetLimit() == 0 || req.GetLimit() > constants.DefaultLimit {
					return constants.DefaultLimit
				}
				return req.GetLimit()
			}(),
			Offset: func() int32 {
				if req.GetOffset() == 0 {
					return constants.DefaultOffset
				}
				return req.GetOffset()
			}(),
			UserID: user.UserID,
		})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch diagnostics: %v", err)
	}
	var responseDiagnostics []*diagnostic.Diagnostic
	for _, diag := range diagnostics {
		responseDiagnostics = append(responseDiagnostics, &diagnostic.Diagnostic{
			DiagnosticId:         diag.ID,
			UserId:               diag.UserID,
			DiagnosticCentreName: diag.DiagnosticCentreName,
			Longitude:            diag.Longitude.Float64,
			Latitude:             diag.Latitude.Float64,
			Address: func() *structpb.Struct {
				var addressStruct structpb.Struct
				if err := addressStruct.UnmarshalJSON(diag.Address); err != nil {
					utils.LogError("Failed to unmarshal address: ", err)
					return nil
				}
				return &addressStruct
			}(),
			Contact: func() *structpb.Struct {
				var contactStruct structpb.Struct
				if err := contactStruct.UnmarshalJSON(diag.Contact); err != nil {
					utils.LogError("Failed to unmarshal contact: ", err)
					return nil
				}
				return &contactStruct
			}(),
			CreatedAt: diag.CreatedAt.Time.String(),
			UpdatedAt: diag.UpdatedAt.Time.String(),
		})
	}
	return &diagnostic.ListDiagnosticsResponse{
		Diagnostics: responseDiagnostics,
	}, nil
}

func (s *DiagnosticService) UpdateDiagnostic(ctx context.Context, req *diagnostic.UpdateDiagnosticRequest) (*diagnostic.UpdateDiagnosticResponse, error) {
	// Validate the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumDIAGNOSTICCENTRE))
	if err != nil {
		utils.LogError("Error validating diagnostic centre: ", err)
		return nil, status.Errorf(codes.PermissionDenied, "Unauthorized: %v", err)
	}
	// Update the diagnostic in the database
	diag, err := s.Repo.UpdateDiagnostic(ctx, db.UpdateDiagnosticParams{
		ID:     req.GetDiagnosticId(),
		UserID: user.UserID,
		DiagnosticCentreName: func() string {
			if req.GetDiagnosticCentreName() == "" {
				oldDiag, err := s.Repo.GetDiagnostic(ctx, req.GetDiagnosticId())
				if err != nil {
					utils.LogError("Failed to fetch existing diagnostic: ", err)
					return ""
				}
				return oldDiag.DiagnosticCentreName
			}
			return req.GetDiagnosticCentreName()
		}(),
		Latitude:  pgtype.Float8{Float64: req.GetLatitude(), Valid: true},
		Longitude: pgtype.Float8{Float64: req.GetLongitude(), Valid: true},
		Address: func() []byte {
			addressBytes, err := json.Marshal(req.GetAddress())
			if err != nil {
				utils.LogError("Failed to marshal address: ", err)
				return nil
			}
			return addressBytes
		}(),
		Contact: func() []byte {
			contactBytes, err := json.Marshal(req.GetContact())
			if err != nil {
				utils.LogError("Failed to marshal contact: ", err)
				return nil
			}
			return contactBytes
		}(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update diagnostic: %v", err)
	}
	return &diagnostic.UpdateDiagnosticResponse{
		DiagnosticId:         diag.ID,
		UserId:               diag.UserID,
		DiagnosticCentreName: diag.DiagnosticCentreName,
		Latitude:             diag.Latitude.Float64,
		Longitude:            diag.Longitude.Float64,
		Address: func() *diagnostic.Address {
			if diag.Address == nil {
				return nil
			}
			var address diagnostic.Address
			if err := json.Unmarshal(diag.Address, &address); err != nil {
				utils.LogError("Failed to unmarshal address: ", err)
				return nil
			}
			return &address
		}(),
		Contact: func() *diagnostic.Contact {
			if diag.Contact == nil {
				return nil
			}
			var contact diagnostic.Contact
			if err := json.Unmarshal(diag.Contact, &contact); err != nil {
				utils.LogError("Failed to unmarshal contact: ", err)
				return nil
			}
			return &contact
		}(),
		CreatedAt: diag.CreatedAt.Time.String(),
		UpdatedAt: diag.UpdatedAt.Time.String(),
	}, nil
}

func (s *DiagnosticService) ListDiagnosticSchedules(ctx context.Context, req *diagnostic.ListDiagnosticSchedulesRequest) (*diagnostic.ListDiagnosticSchedulesResponse, error) {
	// Validate the request
	_, err := middleware.ValidateUser(ctx, string(db.UserEnumDIAGNOSTICCENTRE))
	if err != nil {
		utils.LogError("Error validating diagnostic centre: ", err)
		return nil, status.Errorf(codes.PermissionDenied, "Unauthorized: %v", err)
	}

	// List diagnostic schedules from the database
	diaSchedules, err := s.Repo.ListDiagnosticSchedules(ctx, db.LisDiagnosticSchedulesParams{
		ID: req.DiagnosticId,
		// If Date is not specify
		Column2: func() string {
			if req.GetDate() == "" {
				return time.Now().Format(time.RFC3339)
			}
			return req.GetDate()
		}(),
		// Column2: req.GetDate(),
		// Column3: func() string {
		// 	if req.GetTestType() == "" {
		// 		return ""
		// 	}
		// 	return req.GetTestType()
		// }(),
		// Column3: func() string {
		// 	if req.GetTestStatus() == "" {
		// 		return "SCHEDULED"
		// 	}
		// 	return req.GetTestStatus()
		// }(),
		Limit: func() int32 {
			if req.GetLimit() == 0 || req.GetLimit() > constants.DefaultLimit {
				return constants.DefaultLimit
			}
			return req.GetLimit()
		}(),
		Offset: func() int32 {
			if req.GetOffset() == 0 {
				return constants.DefaultOffset
			}
			return req.GetOffset()
		}(),
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch diagnostic schedules: %v", err)
	}
	var responseSchedules []*diagnostic.GetDiagnosticScheduleResponse
	for _, schedule := range diaSchedules {
		responseSchedules = append(responseSchedules, &diagnostic.GetDiagnosticScheduleResponse{
			ScheduleId:           schedule.ScheduleID,
			DiagnosticCentreName: schedule.DiagnosticCentreName,
			UserId:               schedule.UserID,
			DiagnosticId:         req.GetDiagnosticId(),
			Date:                 schedule.Date.Time.String(),
			Time:                 schedule.Time.Time.String(),
			TestType:             fmt.Sprintf("%v", schedule.TestType),
			Status:               fmt.Sprintf("%v", schedule.Status),
			CreatedAt:            schedule.CreatedAt.Time.String(),
			UpdatedAt:            schedule.UpdatedAt.Time.String(),
		})
	}
	return &diagnostic.ListDiagnosticSchedulesResponse{
		Schedules: responseSchedules,
	}, nil
}
