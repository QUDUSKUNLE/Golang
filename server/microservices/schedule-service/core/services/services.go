package services

import (
	"context"

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
	// Validate the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		utils.LogError("Error validating user: ", err)
		return nil, status.Errorf(codes.PermissionDenied, "Unauthorized: %v", err)
	}
	// N.B Work on validating the user as a registered user
	// Parse timestamps
	date, err := utils.ParseTimestamp(req.GetDate())
	if err != nil {
		utils.LogError("Error parsing date: ", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid date format: %v", err)
	}
	tim, err := utils.ParseTimestamp(req.GetTime())
	if err != nil {
		utils.LogError("Error parsing time: ", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid time format: %v", err)
	}
	// Save the schedule to the database
	response, err := s.Repo.CreateSchedule(ctx, db.CreateScheduleParams{
		UserID:             user.UserID,
		DiagnosticCentreID: req.GetDiagnosticCentreId(),
		Date:               date,
		Time:               tim,
		TestType:           db.ScheduleType(req.GetTestType().Enum().String()),
		Status:             db.ScheduleStatus(req.GetStatus().Enum().String()),
		Notes: func() pgtype.Text {
			var text pgtype.Text
			_ = text.Scan(req.GetNotes()) // Assuming req.GetNotes() is in a valid format
			return text
		}(),
	})
	if err != nil {
		utils.LogError("Error creating schedule: ", err)
		return nil, status.Errorf(codes.Internal, "Database error while creating schedule: %v", err)
	}
	return &schedule.ScheduleResponse{
		ScheduleId: response.ID,
		UserId:     user.UserID,
		Message:    "Schedule created successfully",
		CreatedAt:  response.CreatedAt.Time.String(),
		UpdatedAt:  response.UpdatedAt.Time.String(),
	}, nil
}

func (s *ScheduleService) GetScheduleSession(ctx context.Context, arg *schedule.GetScheduledSessionRequest) (*schedule.GetScheduledSessionResponse, error) {
	// Validate the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		utils.LogError("Error validating user: ", err)
		return nil, status.Errorf(codes.PermissionDenied, "Unauthorized: %v", err)
	}

	// Get the schedule by ID from the database
	response, err := s.Repo.GetScheduleByID(ctx, db.GetScheduleParams{ID: arg.GetScheduleId(), UserID: user.UserID})
	if err != nil {
		utils.LogError("Error getting schedule: ", err)
		return nil, status.Errorf(codes.Internal, "Database error while getting schedule: %v", err)
	}
	return &schedule.GetScheduledSessionResponse{
		ScheduleId:         response.ID,
		UserId:             response.UserID,
		DiagnosticCentreId: response.DiagnosticCentreID,
		Date:               response.Date.Time.String(),
		Time:               response.Time.Time.String(),
		TestType:           string(response.TestType),
		Status:             string(response.Status),
		Notes:              &response.Notes.String,
		CreatedAt:          response.CreatedAt.Time.String(),
		UpdatedAt:          response.UpdatedAt.Time.String(),
	}, nil
}

func (s *ScheduleService) CancelScheduleSession(ctx context.Context, req schedule.CancelScheduledSessionRequest) (*schedule.CancelScheduledSessionResponse, error) {
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
	return &schedule.CancelScheduledSessionResponse{
		ScheduleId: response.ID,
		UserId:     user.UserID,
		Message:    "Schedule cancelled successfully",
	}, nil
}

func (s *ScheduleService) UpdateScheduleSession(ctx context.Context, req *schedule.UpdateScheduledSessionRequest) (*schedule.UpdateScheduledSessionResponse, error) {
	// Validate the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		utils.LogError("Error validating user: ", err)
		// Return an error if the user is not valid
		return nil, err
	}
	// Ensure the user is authorized to access the schedule
	if req.GetUserId() != user.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "You are not authorized to access this schedule")
	}
	// Update the schedule in the database
	response, err := s.Repo.UpdateSchedule(ctx, db.UpdateScheduleParams{
		ID:       req.GetScheduleId(),
		UserID:   user.UserID,
		Date:     parseTimestamp(req.GetDate()),
		Time:     parseTimestamp(req.GetTime()),
		TestType: db.ScheduleType(req.GetTestType()),
		Status:   db.ScheduleStatus(req.GetStatus()),
		Notes: func() pgtype.Text {
			var text pgtype.Text
			_ = text.Scan(req.GetNotes()) // Assuming req.GetNotes() is in a valid format
			return text
		}(),
	})
	if err != nil {
		utils.LogError("Error updating schedule: ", err)
		return nil, err
	}
	return &schedule.UpdateScheduledSessionResponse{
		ScheduleId: response.ID,
		UserId:     response.UserID,
		Message:    "Schedule updated successfully",
		CreatedAt:  response.CreatedAt.Time.String(),
		UpdatedAt:  response.UpdatedAt.Time.String(),
	}, nil
}

func (s *ScheduleService) ListScheduleSessions(ctx context.Context, req *schedule.ListScheduledSessionsRequest) (*schedule.ListScheduledSessionsResponse, error) {
	// Validate the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		utils.LogError("Error validating user: ", err)
		return nil, err
	}
	// NB Get Schedules by user ID
	// Get all schedules for the user from the database
	response, err := s.Repo.GetUserSchedules(ctx, db.GetSchedulesParams{UserID: user.UserID, Limit: constants.DefaultLimit, Offset: constants.DefaultOffset})
	if err != nil {
		utils.LogError("Error getting user schedules: ", err)
		return nil, err
	}
	schedules := make([]*schedule.ScheduledSession, len(response))
	for i, schedul := range response {
		schedules[i] = &schedule.ScheduledSession{
			ScheduleId:         schedul.ID,
			UserId:             schedul.UserID,
			DiagnosticCentreId: schedul.DiagnosticCentreID,
			Date:               schedul.Date.Time.String(),
			Time:               schedul.Time.Time.String(),
			TestType:           string(schedul.TestType),
			Status:             string(schedul.Status),
			Notes:              &schedul.Notes.String,
			CreatedAt:          schedul.CreatedAt.Time.String(),
			UpdatedAt:          schedul.UpdatedAt.Time.String(),
		}
	}
	return &schedule.ListScheduledSessionsResponse{Sessions: schedules}, nil
}

func (s *ScheduleService) ListDiagnosticCentreSchedules(ctx context.Context, req *schedule.ListDiagnosticCentreSchedulesRequest) (*schedule.ListDiagnosticCentreSchedulesResponse, error) {
	// Validate the request
	_, err := middleware.ValidateUser(ctx, string(db.UserEnumDIAGNOSTICCENTRE))
	if err != nil {
		utils.LogError("Error validating user: ", err)
		return nil, err
	}

	// Parse the date if provided
	var parsedDate *pgtype.Timestamptz
	if req.GetDate() != "" {
		date, err := utils.ParseTimestamp(req.GetDate())
		if err != nil {
			utils.LogError("Error parsing date: ", err)
			return nil, status.Errorf(codes.InvalidArgument, "Invalid date format: %v", err)
		}
		parsedDate = &date
	}
	// Convert parsedDate to pgtype.Timestamp if not nil
	var convertedDate pgtype.Timestamp
	if parsedDate != nil {
		convertedDate = pgtype.Timestamp{Time: parsedDate.Time, Valid: parsedDate.Valid}
	}
	// Get all schedules for the diagnostic centre from the database
	response, err := s.Repo.GetSchedulesByDiagnosticCentre(ctx, db.GetSchedulesByDiagnosticCentreParams{Limit: req.GetLimit(), Offset: req.GetOffset(), DiagnosticCentreID: req.GetDiagnosticCentreId(), Column2: db.ScheduleStatus(req.GetStatus().Enum().String()), Column3: convertedDate})
	if err != nil {
		utils.LogError("Error getting diagnostic centre schedules: ", err)
		return nil, err
	}
	schedules := make([]*schedule.ScheduledSession, len(response))
	for i, schedul := range response {
		schedules[i] = &schedule.ScheduledSession{
			ScheduleId:         schedul.ID,
			UserId:             schedul.UserID,
			DiagnosticCentreId: schedul.DiagnosticCentreID,
			Date:               schedul.Date.Time.String(),
			Time:               schedul.Time.Time.String(),
			TestType:           string(schedul.TestType),
			Status:             string(schedul.Status),
			Notes:              &schedul.Notes.String,
			CreatedAt:          schedul.CreatedAt.Time.String(),
			UpdatedAt:          schedul.UpdatedAt.Time.String(),
		}
	}
	return &schedule.ListDiagnosticCentreSchedulesResponse{Sessions: schedules}, nil
}

func (s *ScheduleService) GetScheduleByDiagnosticCentre(ctx context.Context, req *schedule.GetDiagnosticCentreScheduleRequest) (*schedule.GetDiagnosticCentreScheduleResponse, error) {
	// Validate the request
	_, err := middleware.ValidateUser(ctx, string(db.UserEnumDIAGNOSTICCENTRE))
	if err != nil {
		utils.LogError("Error validating user: ", err)
		return nil, err
	}
	// Get the schedule by ID from the database
	response, err := s.Repo.GetScheduleByDiagnosticCentre(ctx, db.GetScheduleByDiagnosticCentreParams{ID: req.GetScheduleId(), DiagnosticCentreID: req.GetDiagnosticCentreId()})
	if err != nil {
		utils.LogError("Error getting diagnostic centre schedule: ", err)
		return nil, err
	}
	return &schedule.GetDiagnosticCentreScheduleResponse{
		ScheduleId:         response.ID,
		UserId:             response.UserID,
		DiagnosticCentreId: response.DiagnosticCentreID,
		Date:               response.Date.Time.String(),
		Time:               response.Time.Time.String(),
		TestType:           string(response.TestType),
		Status:             string(response.Status),
		Notes:              &response.Notes.String,
	}, nil
}
