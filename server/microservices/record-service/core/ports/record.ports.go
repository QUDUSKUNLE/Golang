package ports

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/record-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/record-service/db"
)

type RepositoryPorts interface {
	CreateRecord(ctx context.Context, record domain.RecordDto) (*db.Record, error)
	GetRecord(ctx context.Context, id string) (*db.Record, error)
	GetRecords(ctx context.Context, organizationID string) ([]*db.Record, error)
}

type UseCasePorts interface {
	CreateRecord(ctx context.Context, record domain.RecordDto) (*db.Record, error)
	GetRecord(ctx context.Context, id string) (*db.Record, error)
	GetRecords(ctx context.Context, organizationID string) ([]*db.Record, error)
}
