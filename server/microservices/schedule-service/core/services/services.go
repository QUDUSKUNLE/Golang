package services

import (
	"context"
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
	// Validate the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		utils.LogError("Error validating user: ", err)
		return nil, status.Errorf(codes.PermissionDenied, "Unauthorized: %v", err)
	}
	// N.B Work on validating the user as a registered user
	// Parse timestamps
	date, err := utils.ParseTimestampToPgTimestamptz(req.GetDate())
	if err != nil {
		utils.LogError("Error parsing date: ", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid date format: %v", err)
	}
	tim, err := utils.ParseTimestampToPgTimestamptz(req.GetTime())
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
		Status:             db.ScheduleStatus(req.GetTestStatus().Enum().String()),
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

func (s *ScheduleService) GetScheduleSession(ctx context.Context, arg *schedule.GetScheduleRequest) (*schedule.GetScheduleResponse, error) {
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
		Status:   db.ScheduleStatus(req.GetTestStatus()),
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
	return &schedule.UpdateScheduleResponse{
		ScheduleId: response.ID,
		UserId:     response.UserID,
		Message:    "Schedule updated successfully",
		CreatedAt:  response.CreatedAt.Time.String(),
		UpdatedAt:  response.UpdatedAt.Time.String(),
	}, nil
}

func (s *ScheduleService) ListScheduleSessions(ctx context.Context, req *schedule.ListSchedulesRequest) (*schedule.ListSchedulesResponse, error) {
	// Validate the request
	user, err := middleware.ValidateUser(ctx, string(db.UserEnumUSER))
	if err != nil {
		utils.LogError("Error validating user: ", err)
		return nil, err
	}
	// NB Get Schedules by user ID
	// Get all schedules for the user from the database
	response, err := s.Repo.GetUserSchedules(
		ctx,
		db.GetSchedulesParams{
			UserID: user.UserID,
			Limit: func() int32 {
				// Set a default limit if none is provided or if the limit is greater than 50
				if req.GetLimit() == 0 || req.GetLimit() > 50 {
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
			TestType:           schedule.ScheduleType(schedule.ScheduleType_value[string(schedul.TestType)]),
			TestStatus:         schedule.ScheduleStatus(schedule.ScheduleStatus_value[string(schedul.Status)]),
			Notes:              &schedul.Notes.String,
			CreatedAt:          schedul.CreatedAt.Time.String(),
			UpdatedAt:          schedul.UpdatedAt.Time.String(),
		}
	}
	return &schedule.ListSchedulesResponse{Sessions: schedules}, nil
}

// func (s *ScheduleService) ListDiagnosticCentreSchedules(ctx context.Context, req *schedule.ListDiagnosticCentreSchedulesRequest) (*schedule.ListDiagnosticCentreSchedulesResponse, error) {
// 	// Validate the request
// 	_, err := middleware.ValidateUser(ctx, string(db.UserEnumDIAGNOSTICCENTRE))
// 	if err != nil {
// 		utils.LogError("Error validating user: ", err)
// 		return nil, err
// 	}

// 	// Parse the date if provided
// 	var parsedDate pgtype.Timestamptz
// 	if req.GetDate() == "" {
// 		date, err := utils.ParseTimestampToPgTimestamptz(time.Now().Format(time.RFC3339))
// 		if err != nil {
// 			utils.LogError("Error parsing date: ", err)
// 			return nil, status.Errorf(codes.InvalidArgument, "Invalid date format: %v", err)
// 		}
// 		parsedDate = date
// 	} else {
// 		date, err := utils.ParseTimestampToPgTimestamptz(req.GetDate())
// 		if err != nil {
// 			utils.LogError("Error parsing date: ", err)
// 			return nil, status.Errorf(codes.InvalidArgument, "Invalid date format: %v", err)
// 		}
// 		parsedDate = date
// 	}
// 	// Get all schedules for the diagnostic centre from the database
// 	response, err := s.Repo.GetSchedulesDiagnosticCentreByStatusAndDate(ctx, db.GetSchedulesDiagnosticCentreByStatusAndDateParams{
// 		Limit: func() int32 {
// 			if req.GetLimit() == 0 {
// 				return constants.DefaultLimit
// 			}
// 			return req.GetLimit()
// 		}(),
// 		Offset: func() int32 {
// 			if req.GetOffset() == 0 {
// 				return constants.DefaultOffset
// 			}
// 			return req.GetOffset()
// 		}(),
// 		DiagnosticCentreID: req.GetDiagnosticCentreId(),
// 		Status: func() db.ScheduleStatus {
// 			if req.GetTestStatus() == schedule.TestStatus_TEST_STATUS_UNSPECIFIED {
// 				return db.ScheduleStatus("SCHEDULED")
// 			}
// 			return db.ScheduleStatus(req.GetTestStatus().Enum().String())
// 		}(),
// 		Date: parsedDate,
// 	})
// 	if err != nil {
// 		utils.LogError("Error getting diagnostic centre schedules: ", err)
// 		return nil, err
// 	}
// 	schedules := make([]*schedule.ScheduledSession, len(response))
// 	for i, schedul := range response {
// 		schedules[i] = &schedule.ScheduledSession{
// 			ScheduleId:         schedul.ID,
// 			UserId:             schedul.UserID,
// 			DiagnosticCentreId: schedul.DiagnosticCentreID,
// 			Date:               schedul.Date.Time.String(),
// 			Time:               schedul.Time.Time.String(),
// 			TestType:           schedule.TestType(schedule.TestType_value[string(schedul.TestType)]),
// 			TestStatus:         schedule.TestStatus(schedule.TestStatus_value[string(schedul.Status)]),
// 			Notes:              &schedul.Notes.String,
// 			CreatedAt:          schedul.CreatedAt.Time.String(),
// 			UpdatedAt:          schedul.UpdatedAt.Time.String(),
// 		}
// 	}
// 	return &schedule.ListDiagnosticCentreSchedulesResponse{Sessions: schedules}, nil
// }

// func (s *ScheduleService) GetScheduleByDiagnosticCentre(ctx context.Context, req *schedule.GetDiagnosticCentreScheduleRequest) (*schedule.GetDiagnosticCentreScheduleResponse, error) {
// 	// Validate the request
// 	_, err := middleware.ValidateUser(ctx, string(db.UserEnumDIAGNOSTICCENTRE))
// 	if err != nil {
// 		utils.LogError("Error validating user: ", err)
// 		return nil, err
// 	}
// 	// Get the schedule by ID from the database
// 	response, err := s.Repo.GetScheduleByDiagnosticCentre(ctx, db.GetScheduleByDiagnosticCentreParams{ID: req.GetScheduleId(), DiagnosticCentreID: req.GetDiagnosticCentreId()})
// 	if err != nil {
// 		utils.LogError("Error getting diagnostic centre schedule: ", err)
// 		return nil, err
// 	}
// 	return &schedule.GetDiagnosticCentreScheduleResponse{
// 		ScheduleId:         response.ID,
// 		UserId:             response.UserID,
// 		DiagnosticCentreId: response.DiagnosticCentreID,
// 		Date:               response.Date.Time.String(),
// 		Time:               response.Time.Time.String(),
// 		TestType:           schedule.TestType(schedule.TestType_value[string(response.TestType)]),
// 		TestStatus:         schedule.TestStatus(schedule.TestStatus_value[string(response.Status)]),
// 		Notes:              &response.Notes.String,
// 	}, nil
// }
