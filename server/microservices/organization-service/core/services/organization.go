package services

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/organization-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/organization-service/db"
)

type Repository struct {
	database *db.Queries
}

// CreateOrganization implements ports.RepositoryPorts.
func (r *Repository) CreateOrganization(ctx context.Context, user domain.OrganizationDto) (*db.Organization, error) {
	return r.database.CreateOrganization(ctx, user.UserID)
}

// GetOrganization implements ports.RepositoryPorts.
func (r *Repository) GetOrganization(ctx context.Context, id string) (*db.Organization, error) {
	return r.database.GetOrganization(ctx, id)
}

func NewRepository(database *db.Queries) ports.RepositoryPorts {
	return &Repository{database: database}
}
