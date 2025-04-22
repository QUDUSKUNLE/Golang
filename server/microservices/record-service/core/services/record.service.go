package services

import (
	"context"
	"strings"

	"github.com/QUDUSKUNLE/microservices/gateway/db"
	"github.com/QUDUSKUNLE/microservices/record-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/record-service/core/ports"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	database *db.Queries
}

// SearchRecordByNin implements ports.RepositoryPorts.
func (r *Repository) SearchRecordByNin(ctx context.Context, searchRecord domain.GetRecordByNinDto) ([]*db.SearchRecordByNinRow, error) {
	nin := pgtype.Text{String: searchRecord.Nin, Valid: true}
	if strings.TrimSpace(searchRecord.ScanTitle) != "" {
		return r.database.SearchRecordByNin(ctx, nin)
	}
	result, err := r.database.SearchRecordByNinAndScanTitle(ctx, db.SearchRecordByNinAndScanTitleParams{Nin: nin, ScanTitle: "%" + strings.TrimSpace(searchRecord.ScanTitle) + "%"})
	if err != nil {
		return nil, err
	}
	actualResult := make([]*db.SearchRecordByNinRow, 0)
	for _, r := range result {
		actualResult = append(actualResult, &db.SearchRecordByNinRow{
			ID:             r.ID,
			OrganizationID: r.OrganizationID,
			UserID:         r.UserID,
			Record:         r.Record,
			ScanTitle:      r.ScanTitle,
			CreatedAt:      r.CreatedAt,
			UpdatedAt:      r.UpdatedAt,
		})
	}
	return actualResult, nil
}

// SearchRecord implements ports.RepositoryPorts.
func (r *Repository) SearchRecord(ctx context.Context, searchRecord domain.GetRecordDto) ([]*db.Record, error) {
	if strings.TrimSpace(searchRecord.ScanTitle) != "" {
		return r.database.GetRecordsByUserAndScanTitle(ctx, db.GetRecordsByUserAndScanTitleParams{UserID: searchRecord.UserID, ScanTitle: "%" + strings.TrimSpace(searchRecord.ScanTitle) + "%"})
	}
	return r.database.GetRecordsByUser(ctx, searchRecord.UserID)
}

// UploadRecord implements ports.RepositoryPorts.
func (r *Repository) UploadRecord(ctx context.Context, record domain.UploadDto) (*db.Upload, error) {
	return r.database.UploadRecord(ctx, db.UploadRecordParams{UserID: record.UserID, OrganizationID: record.OrganizationID, ScanTitle: record.ScanTitle})
}

// GetRecords implements ports.RepositoryPorts.
func (r *Repository) GetRecords(ctx context.Context, organizationID string) ([]*db.Record, error) {
	return r.database.GetRecords(ctx, organizationID)
}

// GetRecord implements ports.RepositoryPorts.
func (r *Repository) GetRecord(ctx context.Context, record string) (*db.Record, error) {
	return r.database.GetRecord(ctx, record)
}

// CreateRecord implements ports.RepositoryPorts.
func (r *Repository) CreateRecord(ctx context.Context, record domain.RecordDto) (*db.Record, error) {
	return r.database.CreateRecord(ctx, db.CreateRecordParams{OrganizationID: record.OrganizationID, UserID: record.UserID, Record: record.Record, ScanTitle: record.ScanTitle})
}

func NewRepository(database *db.Queries) ports.RepositoryPorts {
	return &Repository{
		database: database,
	}
}
