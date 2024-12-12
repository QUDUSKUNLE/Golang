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

// CreateRecord implements ports.UseCasePorts.
func (u *UseCase) CreateRecord(ctx context.Context, record domain.RecordDto) (interface{}, error) {
	return u.usecase.CreateRecord(ctx, record)
}

func InitializeRecordService(db *db.Queries) ports.UseCasePorts {
	records := services.NewRepository(db)
	return &UseCase{usecase: records}
}
