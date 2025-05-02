package repository

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/db"
)

type ScheduleRepository struct {
	database *db.Queries
}

func (s *ScheduleRepository) CreateSchedule(ctx context.Context, arg db.CreateScheduleParams) (*db.DiagnosticSchedule, error) {
	// Implementation for creating a schedule
	schedule, err := s.database.CreateSchedule(ctx, arg)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

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

func (s *ScheduleRepository) GetSchedulesByStatus(ctx context.Context, arg db.GetSchedulesByStatusParams) ([]*db.DiagnosticSchedule, error) {
	// Implementation for getting schedules by status
	schedules, err := s.database.GetSchedulesByStatus(ctx, arg)
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

func NewScheduleRepository(database *db.Queries) *ScheduleRepository {
	return &ScheduleRepository{
		database: database,
	}
}
