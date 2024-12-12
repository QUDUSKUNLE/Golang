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

// CreateRecord implements ports.RepositoryPorts.
func (r *Repository) CreateRecord(ctx context.Context, record domain.RecordDto) (interface{}, error) {
	return r.database.CreateRecord(ctx, db.CreateRecordParams{OrganizationID: record.OrganizationID, UserID: record.UserID, Record: record.Record})
}

func NewRepository(database *db.Queries) ports.RepositoryPorts {
	return &Repository{
		database: database,
	}
}
