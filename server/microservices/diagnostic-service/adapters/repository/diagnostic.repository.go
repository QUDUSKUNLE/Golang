package repository

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/db"
)

type DiagnosticRepository struct {
	database *db.Queries
}

func (d *DiagnosticRepository) CreateDiagnostic(ctx context.Context, arg db.CreateDiagnosticParams) (*db.Diagnostic, error) {
	// Implementation for creating a diagnostic
	diagnostic, err := d.database.CreateDiagnostic(ctx, arg)
	if err != nil {
		return nil, err
	}
	return diagnostic, nil
}

func (d *DiagnosticRepository) GetDiagnostic(ctx context.Context, id string) (*db.Diagnostic, error) {
	// Implementation for getting a diagnostic by ID
	diagnostic, err := d.database.GetDiagnostic(ctx, id)
	if err != nil {
		return nil, err
	}
	return diagnostic, nil
}

func (d *DiagnosticRepository) GetAllDiagnostics(ctx context.Context, arg db.GetAllDiagnosticsParams) ([]*db.Diagnostic, error) {
	// Implementation for getting all diagnostics
	diagnostics, err := d.database.GetAllDiagnostics(ctx, arg)
	if err != nil {
		return nil, err
	}
	return diagnostics, nil
}

func (d *DiagnosticRepository) CancelDiagnostic(ctx context.Context, diagnostic_id string) (*db.Diagnostic, error) {
	// Implementation for deleting a diagnostic
	diagnostic, err := d.database.DeleteDiagnostic(ctx, diagnostic_id)
	if err != nil {
		return nil, err
	}
	return diagnostic, nil
}

func (d *DiagnosticRepository) ListDiagnostics(ctx context.Context, arg db.ListDiagnosticsParams) ([]*db.Diagnostic, error) {
	// Implementation for listing diagnostics
	diagnostics, err := d.database.ListDiagnostics(ctx, arg)
	if err != nil {
		return nil, err
	}
	return diagnostics, nil
}

func (d *DiagnosticRepository) SearchDiagnostics(ctx context.Context, arg db.SearchDiagnosticsParams) ([]*db.Diagnostic, error) {
	// Implementation for searching diagnostics
	diagnostics, err := d.database.SearchDiagnostics(ctx, arg)
	if err != nil {
		return nil, err
	}
	return diagnostics, nil
}

func (d *DiagnosticRepository) UpdateDiagnostic(ctx context.Context, arg db.UpdateDiagnosticParams) (*db.Diagnostic, error) {
	// Implementation for updating a diagnostic
	diagnostic, err := d.database.UpdateDiagnostic(ctx, arg)
	if err != nil {
		return nil, err
	}
	return diagnostic, nil
}

func (d *DiagnosticRepository) ListDiagnosticSchedules(ctx context.Context, arg db.LisDiagnosticSchedulesParams) ([]*db.LisDiagnosticSchedulesRow, error) {
	// Implementation for listing diagnostic schedules
	schedules, err := d.database.LisDiagnosticSchedules(ctx, arg)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func (d *DiagnosticRepository) GetDiagnosticSchedule(ctx context.Context, arg db.GetDiagnosticScheduleParams) (*db.GetDiagnosticScheduleRow, error) {
	// Implementation for getting a diagnostic schedule
	schedule, err := d.database.GetDiagnosticSchedule(ctx, arg)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func NewDiagnosticRepository(database *db.Queries) *DiagnosticRepository {
	return &DiagnosticRepository{
		database: database,
	}
}
