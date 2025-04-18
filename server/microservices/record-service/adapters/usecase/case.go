package usecase

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/record-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/record-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/record-service/core/services"
	"github.com/QUDUSKUNLE/microservices/record-service/db"
)

type UseCase struct {
	usecase ports.RepositoryPorts
}

// SearchRecord implements ports.UseCasePorts.
func (u *UseCase) SearchRecord(ctx context.Context, searchRecord domain.GetRecordDto) ([]*db.Record, error) {
	panic("unimplemented")
}

// UploadRecord implements ports.UseCasePorts.
func (u *UseCase) UploadRecord(ctx context.Context, record domain.UploadDto) (*db.Upload, error) {
	return u.usecase.UploadRecord(ctx, record)
}

// GetRecords implements ports.UseCasePorts.
func (u *UseCase) GetRecords(ctx context.Context, organizationID string) ([]*db.Record, error) {
	return u.usecase.GetRecords(ctx, organizationID)
}

// GetRecord implements ports.UseCasePorts.
func (u *UseCase) GetRecord(ctx context.Context, id string) (*db.Record, error) {
	return u.usecase.GetRecord(ctx, id)
}

// CreateRecord implements ports.UseCasePorts.
func (u *UseCase) CreateRecord(ctx context.Context, record domain.RecordDto) (*db.Record, error) {
	return u.usecase.CreateRecord(ctx, record)
}

func InitializeRecordService(db *db.Queries) ports.UseCasePorts {
	records := services.NewRepository(db)
	return &UseCase{usecase: records}
}
