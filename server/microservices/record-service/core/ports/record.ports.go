package ports

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/gateway/db"
	"github.com/QUDUSKUNLE/microservices/record-service/core/domain"
)

type RepositoryPorts interface {
	CreateRecord(ctx context.Context, record domain.RecordDto) (*db.Record, error)
	GetRecord(ctx context.Context, id string) (*db.Record, error)
	SearchRecord(ctx context.Context, searchRecord domain.GetRecordDto) ([]*db.Record, error)
	SearchRecordByNin(ctx context.Context, searchRecord domain.GetRecordByNinDto) ([]*db.SearchRecordByNinRow, error)
	GetRecords(ctx context.Context, organizationID string) ([]*db.Record, error)
	UploadRecord(ctx context.Context, record domain.UploadDto) (*db.Upload, error)
}

type UseCasePorts interface {
	CreateRecord(ctx context.Context, record domain.RecordDto) (*db.Record, error)
	GetRecord(ctx context.Context, id string) (*db.Record, error)
	SearchRecord(ctx context.Context, searchRecord domain.GetRecordDto) ([]*db.Record, error)
	SearchRecordByNin(ctx context.Context, searchRecord domain.GetRecordByNinDto) ([]*db.SearchRecordByNinRow, error)
	GetRecords(ctx context.Context, organizationID string) ([]*db.Record, error)
	UploadRecord(ctx context.Context, record domain.UploadDto) (*db.Upload, error)
}
