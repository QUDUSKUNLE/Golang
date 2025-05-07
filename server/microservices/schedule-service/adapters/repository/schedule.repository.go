package repository

import (
	"context"
	"fmt"

	"github.com/QUDUSKUNLE/microservices/shared/db"
)

// ScheduleRepository provides methods to interact with the diagnostic_schedules table.
type ScheduleRepository struct {
	database *db.Queries
}

// Query by schedulers

// CreateSchedule creates a new diagnostic schedule in the database.
func (s *ScheduleRepository) CreateSchedule(ctx context.Context, arg db.CreateScheduleParams) (*db.DiagnosticSchedule, error) {
	// Implementation for creating a schedule
	schedule, err := s.database.CreateSchedule(ctx, arg)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

// GetScheduleByID retrieves a diagnostic schedule by its ID and user ID.
func (s *ScheduleRepository) GetScheduleByID(ctx context.Context, arg db.GetScheduleParams) (*db.DiagnosticSchedule, error) {
	// Implementation for getting a schedule by ID
	schedule, err := s.database.GetSchedule(ctx, arg)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (s *ScheduleRepository) DeleteSchedule(ctx context.Context, arg db.CancelScheduleParams) (*db.DiagnosticSchedule, error) {
	// Implementation for deleting a schedule
	schedule, err := s.database.CancelSchedule(ctx, arg)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (s *ScheduleRepository) UpdateSchedule(ctx context.Context, arg db.UpdateScheduleParams) (*db.DiagnosticSchedule, error) {
	// Implementation for updating a schedule
	schedule, err := s.database.UpdateSchedule(ctx, arg)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (s *ScheduleRepository) GetSchedulesByCentre(ctx context.Context, arg db.GetSchedulesByCentreParams) ([]*db.DiagnosticSchedule, error) {
	// Implementation for getting schedules by centre
	schedules, err := s.database.GetSchedulesByCentre(ctx, arg)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func (s *ScheduleRepository) GetUserSchedules(ctx context.Context, arg db.GetSchedulesParams) ([]*db.DiagnosticSchedule, error) {
	// Implementation for getting all schedules
	schedules, err := s.database.GetSchedules(ctx, arg)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

// Query by diagnostic centre
func (s *ScheduleRepository) GetScheduleByDiagnosticCentre(ctx context.Context, arg db.GetScheduleByDiagnosticCentreParams) (*db.DiagnosticSchedule, error) {
	// Implementation for getting a schedule by diagnostic centre
	schedule, err := s.database.GetScheduleByDiagnosticCentre(ctx, arg)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (s *ScheduleRepository) GetSchedulesByDiagnosticCentre(ctx context.Context, arg db.GetSchedulesByDiagnosticCentreParams) ([]*db.DiagnosticSchedule, error) {
	// Implementation for getting schedules by diagnostic centre
	schedule, err := s.database.GetSchedulesByDiagnosticCentre(ctx, arg)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (s *ScheduleRepository) GetSchedulesDiagnosticCentreByStatusAndDate(ctx context.Context, arg db.GetSchedulesDiagnosticCentreByStatusAndDateParams) ([]*db.DiagnosticSchedule, error) {
	// Implementation for getting schedules by status and date
	fmt.Println("arg", arg)
	schedules, err := s.database.GetSchedulesDiagnosticCentreByStatusAndDate(ctx, arg)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func NewScheduleRepository(database *db.Queries) *ScheduleRepository {
	return &ScheduleRepository{
		database: database,
	}
}
