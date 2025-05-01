package repository

import (
	"context"
	"github.com/QUDUSKUNLE/microservices/shared/db"
)

type DiagnosticRepository struct {
	database *db.Queries
}

func (d *DiagnosticRepository) CreateDiagnostic(ctx context.Context, userID string) (*db.Diagnostic, error) {
	// Implementation for creating a diagnostic
	diagnostic, err := d.database.CreateDiagnostic(ctx, userID)
	if err != nil {
		return nil, err
	}
	return diagnostic, nil
}

func NewDiagnosticRepository(database *db.Queries) *DiagnosticRepository {
	return &DiagnosticRepository{
		database: database,
	}
}
