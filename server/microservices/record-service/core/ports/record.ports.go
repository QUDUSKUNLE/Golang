package ports

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/record-service/core/domain"
)

type RepositoryPorts interface {
	CreateRecord(ctx context.Context, record domain.RecordDto) (interface{}, error)
}

type UseCasePorts interface {
	CreateRecord(ctx context.Context, record domain.RecordDto) (interface{}, error)
}
