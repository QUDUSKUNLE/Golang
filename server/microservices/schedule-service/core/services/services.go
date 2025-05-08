package services

import (
	"context"
	"fmt"

	// "time"

	"github.com/QUDUSKUNLE/microservices/shared/constants"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/middleware"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/schedule"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Users can schedule appointments for diagnostic tests
func (s *ScheduleService) CreateScheduleSession(ctx context.Context, req *schedule.ScheduleRequest) (*schedule.ScheduleResponse, error) {
	// Validate the user making the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		utils.LogError("User validation failed: ", err)
		return nil, status.Errorf(codes.PermissionDenied, "Unauthorized: %v", err)
	}

	// Parse and validate the date
	date, err := utils.ParseTimestampToPgTimestamptz(req.GetDate())
	if err != nil {
		utils.LogError("Invalid date format: ", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid date format: %v", err)
	}

	// Parse and validate the time
	time, err := utils.ParseTimestampToPgTimestamptz(req.GetTime())
	if err != nil {
		utils.LogError("Invalid time format: ", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid time format: %v", err)
	}

	// Prepare the notes field
	notes := pgtype.Text{}
	if err := notes.Scan(req.GetNotes()); err != nil {
		utils.LogError("Error scanning notes: ", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid notes format: %v", err)
	}
	if db.ScheduleStatus(req.GetTestStatus().Enum().String()) != db.ScheduleStatusSCHEDULED {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid schedule status: %v", req.GetTestStatus().Enum().String())
	}

	fmt.Println(db.ScheduleStatus(req.GetTestStatus().Enum().String()), "status")
	// Save the schedule to the database
	response, err := s.Repo.CreateSchedule(ctx, db.CreateScheduleParams{
		UserID:             user.UserID,
		DiagnosticCentreID: req.GetDiagnosticCentreId(),
		Date:               date,
		Time:               time,
		TestType:           db.ScheduleType(req.GetTestType().Enum().String()),
		Status:             db.ScheduleStatus(req.GetTestStatus().Enum().String()),
		Notes:              notes,
	})
	if err != nil {
		utils.LogError("Database error while creating schedule: ", err)
		return nil, status.Errorf(codes.Internal, "Failed to create schedule: %v", err)
	}

	// Return the response
	return &schedule.ScheduleResponse{
		ScheduleId: response.ID,
		UserId:     user.UserID,
		Message:    "Schedule created successfully",
		CreatedAt:  response.CreatedAt.Time.String(),
		UpdatedAt:  response.UpdatedAt.Time.String(),
	}, nil
}

func (s *ScheduleService) GetScheduleSession(ctx context.Context, req *schedule.GetScheduleRequest) (*schedule.GetScheduleResponse, error) {
	// Validate the user making the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		utils.LogError("User validation failed: ", err)
		return nil, status.Errorf(codes.PermissionDenied, "Unauthorized: %v", err)
	}

	// Ensure the user is authorized to access the schedule
	if req.GetUserId() != user.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "You are not authorized to access this schedule")
	}

	// Retrieve the schedule by ID from the database
	response, err := s.Repo.GetScheduleByID(ctx, db.GetScheduleParams{
		ID:     req.GetScheduleId(),
		UserID: user.UserID,
	})
	if err != nil {
		utils.LogError("Error retrieving schedule: ", err)
		return nil, status.Errorf(codes.Internal, "Failed to retrieve schedule: %v", err)
	}

	// Construct and return the response
	return &schedule.GetScheduleResponse{
		ScheduleId:         response.ID,
		UserId:             response.UserID,
		DiagnosticCentreId: response.DiagnosticCentreID,
		Date:               response.Date.Time.String(),
		Time:               response.Time.Time.String(),
		TestType:           schedule.ScheduleType(schedule.ScheduleType_value[string(response.TestType)]),
		TestStatus:         schedule.ScheduleStatus(schedule.ScheduleStatus_value[string(response.Status)]),
		Notes:              &response.Notes.String,
		CreatedAt:          response.CreatedAt.Time.String(),
		UpdatedAt:          response.UpdatedAt.Time.String(),
	}, nil
}

func (s *ScheduleService) CancelScheduleSession(ctx context.Context, req schedule.CancelScheduleRequest) (*schedule.CancelScheduleResponse, error) {
	// Validate the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		utils.LogError("Error validating user: ", err)
		return nil, err
	}
	// Ensure the user is authorized to access the schedule
	if req.GetUserId() != user.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "You are not authorized to access this schedule")
	}
	// Delete the schedule by ID from the database
	response, err := s.Repo.DeleteSchedule(ctx, db.CancelScheduleParams{ID: req.GetScheduleId(), UserID: user.UserID})
	if err != nil {
		utils.LogError("Error deleting schedule: ", err)
		return nil, err
	}
	return &schedule.CancelScheduleResponse{
		ScheduleId: response.ID,
		UserId:     user.UserID,
		Message:    "Schedule cancelled successfully",
	}, nil
}

func (s *ScheduleService) UpdateScheduleSession(ctx context.Context, req *schedule.UpdateScheduleRequest) (*schedule.UpdateScheduleResponse, error) {
	// Validate the user making the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		utils.LogError("User validation failed: ", err)
		return nil, status.Errorf(codes.PermissionDenied, "Unauthorized: %v", err)
	}

	// Ensure the user is authorized to update the schedule
	if req.GetUserId() != user.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "You are not authorized to update this schedule")
	}

	// Parse and validate the date
	date, err := utils.ParseTimestampToPgTimestamptz(req.GetDate())
	if err != nil {
		utils.LogError("Invalid date format: ", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid date format: %v", err)
	}

	// Parse and validate the time
	time, err := utils.ParseTimestampToPgTimestamptz(req.GetTime())
	if err != nil {
		utils.LogError("Invalid time format: ", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid time format: %v", err)
	}

	// Prepare the notes field
	notes := pgtype.Text{}
	if err := notes.Scan(req.GetNotes()); err != nil {
		utils.LogError("Error scanning notes: ", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid notes format: %v", err)
	}

	// Update the schedule in the database
	response, err := s.Repo.UpdateSchedule(ctx, db.UpdateScheduleParams{
		ID:       req.GetScheduleId(),
		UserID:   user.UserID,
		Date:     date,
		Time:     time,
		TestType: db.ScheduleType(req.GetTestType().Enum().String()),
		Status:   db.ScheduleStatus(req.GetTestStatus().Enum().String()),
		Notes:    notes,
	})
	if err != nil {
		utils.LogError("Error updating schedule: ", err)
		return nil, status.Errorf(codes.Internal, "Failed to update schedule: %v", err)
	}

	// Return the response
	return &schedule.UpdateScheduleResponse{
		ScheduleId: response.ID,
		UserId:     response.UserID,
		Message:    "Schedule updated successfully",
		CreatedAt:  response.CreatedAt.Time.String(),
		UpdatedAt:  response.UpdatedAt.Time.String(),
	}, nil
}

func (s *ScheduleService) ListScheduleSessions(ctx context.Context, req *schedule.ListSchedulesRequest) (*schedule.ListSchedulesResponse, error) {
	// Validate the user making the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		utils.LogError("User validation failed: ", err)
		return nil, status.Errorf(codes.PermissionDenied, "Unauthorized: %v", err)
	}

	// Set default limit and offset if not provided or invalid
	limit := req.GetLimit()
	if limit == 0 || limit > constants.MaxLimit {
		limit = constants.DefaultLimit
	}

	offset := req.GetOffset()
	if offset < 0 {
		offset = constants.DefaultOffset
	}

	// Retrieve schedules for the user from the database
	response, err := s.Repo.GetUserSchedules(ctx, db.GetSchedulesParams{
		UserID: user.UserID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		utils.LogError("Error retrieving user schedules: ", err)
		return nil, status.Errorf(codes.Internal, "Failed to retrieve schedules: %v", err)
	}

	// Map database response to API response
	schedules := make([]*schedule.ScheduledSession, len(response))
	for i, schedul := range response {
		schedules[i] = &schedule.ScheduledSession{
			ScheduleId:         schedul.ID,
			UserId:             schedul.UserID,
			DiagnosticCentreId: schedul.DiagnosticCentreID,
			Date:               schedul.Date.Time.String(),
			Time:               schedul.Time.Time.String(),
			TestType:           schedule.ScheduleType(schedule.ScheduleType_value[string(schedul.TestType)]),
			TestStatus:         schedule.ScheduleStatus(schedule.ScheduleStatus_value[string(schedul.Status)]),
			Notes:              &schedul.Notes.String,
			CreatedAt:          schedul.CreatedAt.Time.String(),
			UpdatedAt:          schedul.UpdatedAt.Time.String(),
		}
	}

	// Return the list of schedules
	return &schedule.ListSchedulesResponse{Sessions: schedules}, nil
}
