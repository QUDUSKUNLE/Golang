package services

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/record-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/record-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/record-service/adapters/repository"
	"github.com/QUDUSKUNLE/microservices/shared/db"
)

type RecordCase struct {
	usecase ports.RepositoryPorts
}

// SearchRecordByNin implements ports.UseCasePorts.
func (u *RecordCase) SearchRecordByNin(ctx context.Context, searchRecord domain.GetRecordByNinDto) ([]*db.SearchRecordByNinRow, error) {
	return u.usecase.SearchRecordByNin(ctx, searchRecord)
}

// SearchRecord implements ports.UseCasePorts.
func (u *RecordCase) SearchRecord(ctx context.Context, searchRecord domain.GetRecordDto) ([]*db.Record, error) {
	return u.usecase.SearchRecord(ctx, searchRecord)
}

// UploadRecord implements ports.UseCasePorts.
func (u *RecordCase) UploadRecord(ctx context.Context, record domain.UploadDto) (*db.Upload, error) {
	return u.usecase.UploadRecord(ctx, record)
}

// GetRecords implements ports.UseCasePorts.
func (u *RecordCase) GetRecords(ctx context.Context, organizationID string) ([]*db.Record, error) {
	return u.usecase.GetRecords(ctx, organizationID)
}

// GetRecord implements ports.UseCasePorts.
func (u *RecordCase) GetRecord(ctx context.Context, id string) (*db.Record, error) {
	return u.usecase.GetRecord(ctx, id)
}

// CreateRecord implements ports.UseCasePorts.
func (u *RecordCase) CreateRecord(ctx context.Context, record domain.RecordDto) (*db.Record, error) {
	return u.usecase.CreateRecord(ctx, record)
}

func InitializeRecordService(db *db.Queries) ports.RecordPorts {
	records := repository.NewRepository(db)
	return &RecordCase{usecase: records}
}
