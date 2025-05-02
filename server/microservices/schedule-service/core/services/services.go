package services

import (
	"context"
	"time"

	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/middleware"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/schedule"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ScheduleService) CreateSchedule(ctx context.Context, req *schedule.ScheduleRequest) (*schedule.ScheduleResponse, error) {
	// Validate the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		return nil, err
	}
	// Save the schedule to the database
	response, err := s.Repo.CreateSchedule(ctx, db.CreateScheduleParams{
		UserID:             user.UserID,
		DiagnosticCentreID: req.GetDiagnosticCentreId(),
		Date: func() pgtype.Timestamptz {
			parsedTime, err := time.Parse(time.RFC3339, req.GetDate())
			if err != nil {
				return pgtype.Timestamptz{Valid: false} // Handle invalid date format
			}
			return pgtype.Timestamptz{Time: parsedTime, Valid: true}
		}(),
		Time: func() pgtype.Timestamptz {
			parsedTime, err := time.Parse(time.RFC3339, req.GetDate())
			if err != nil {
				return pgtype.Timestamptz{Valid: false} // Handle invalid date format
			}
			return pgtype.Timestamptz{Time: parsedTime, Valid: true}
		}(),
		TestType: db.ScheduleType(req.GetTestType().Enum().String()),
		Status:   db.ScheduleStatus(req.GetStatus().Enum().String()),
		Notes: func() pgtype.Text {
			var text pgtype.Text
			_ = text.Scan(req.GetNotes()) // Assuming req.GetNotes() is in a valid format
			return text
		}(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Database error while creating schedule: %v", err)
	}
	return &schedule.ScheduleResponse{
		ScheduleId: response.ID,
		UserId:     user.UserID,
		Message:    "Schedule created successfully",
	}, nil
}

func (s *ScheduleService) GetScheduleSession(ctx context.Context, arg *schedule.GetScheduledSessionRequest) (*schedule.GetScheduledSessionResponse, error) {
	// Validate the request
	_, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		return nil, err
	}
	// Get the schedule by ID from the database
	response, err := s.Repo.GetScheduleByID(ctx, arg.GetScheduleId())
	if err != nil {
		return nil, err
	}
	return &schedule.GetScheduledSessionResponse{
		ScheduleId:         response.ID,
		UserId:             response.UserID,
		DiagnosticCentreId: response.DiagnosticCentreID,
		Date:               response.Date.Time.String(),
		Time:               response.Time.Time.String(),
		TestType:           string(response.TestType),
		Status:             string(response.Status),
		Notes:              response.Notes.String,
		CreatedAt:          response.CreatedAt.Time.String(),
		UpdatedAt:          response.UpdatedAt.Time.String(),
	}, nil
}

func (s *ScheduleService) DeleteSchedule(ctx context.Context, req schedule.CancelScheduledSessionRequest) (*schedule.CancelScheduledSessionResponse, error) {
	// Validate the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		return nil, err
	}
	// Delete the schedule by ID from the database
	response, err := s.Repo.DeleteSchedule(ctx, db.CancelScheduleParams{ID: req.GetScheduleId(), UserID: user.UserID})
	if err != nil {
		return nil, err
	}
	return &schedule.CancelScheduledSessionResponse{
		ScheduleId: response.ID,
		UserId:     user.UserID,
		Message:    "Schedule cancelled successfully",
	}, nil
}

func (s *ScheduleService) UpdateSchedule(ctx context.Context, req *schedule.UpdateScheduledSessionRequest) (*schedule.UpdateScheduledSessionResponse, error) {
	// Validate the request
	_, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		return nil, err
	}
	// Update the schedule in the database
	response, err := s.Repo.UpdateSchedule(ctx, db.UpdateScheduleParams{
		ID:       req.GetScheduleId(),
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

func (s *ScheduleService) GetUserSchedules(ctx context.Context, req *schedule.ListScheduledSessionsRequest) (*schedule.ListScheduledSessionsResponse, error) {
	// Validate the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		return nil, err
	}
	// Get all schedules for the user from the database
	response, err := s.Repo.GetUserSchedules(ctx, db.GetSchedulesParams{UserID: user.UserID})
	if err != nil {
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
			Notes:              schedul.Notes.String,
			CreatedAt:          schedul.CreatedAt.Time.String(),
			UpdatedAt:          schedul.UpdatedAt.Time.String(),
		}
	}
	return &schedule.ListScheduledSessionsResponse{Sessions: schedules}, nil
}
