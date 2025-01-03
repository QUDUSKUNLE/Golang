package services

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/record-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/record-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/record-service/db"
)

type Repository struct {
	database *db.Queries
}

// UploadRecord implements ports.RepositoryPorts.
func (r *Repository) UploadRecord(ctx context.Context, record domain.UploadDto) (*db.Upload, error) {
	return r.database.UploadRecord(ctx, db.UploadRecordParams{UserID: record.UserID, OrganizationID: record.OrganizationID, FileName: record.FileName, ScanTitle: record.ScanTitle})
}

// GetRecords implements ports.RepositoryPorts.
func (r *Repository) GetRecords(ctx context.Context, organizationID string) ([]*db.Record, error) {
	return r.database.GetRecords(ctx, organizationID)
}

// GetRecord implements ports.RepositoryPorts.
func (r *Repository) GetRecord(ctx context.Context, id string) (*db.Record, error) {
	return r.database.GetRecord(ctx, id)
}

// CreateRecord implements ports.RepositoryPorts.
func (r *Repository) CreateRecord(ctx context.Context, record domain.RecordDto) (*db.Record, error) {
	return r.database.CreateRecord(ctx, db.CreateRecordParams{OrganizationID: record.OrganizationID, UserID: record.UserID, Record: record.Record})
}

func NewRepository(database *db.Queries) ports.RepositoryPorts {
	return &Repository{
		database: database,
	}
}
