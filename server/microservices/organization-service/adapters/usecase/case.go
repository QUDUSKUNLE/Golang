package usecase

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/organization-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/services"
	"github.com/QUDUSKUNLE/microservices/organization-service/internal/db"
)

type UseCase struct {
	usecase ports.RepositoryPorts
}

// CreateOrganization implements ports.UseCasePorts.
func (u *UseCase) CreateOrganization(ctx context.Context, user domain.OrganizationDto) (*db.Organization, error) {
	return u.usecase.CreateOrganization(ctx, user)
}

// GetOrganization implements ports.UseCasePorts.
func (u *UseCase) GetOrganization(ctx context.Context, id string) (*db.Organization, error) {
	return u.usecase.GetOrganization(ctx, id)
}

func InitOrganizationServer(db *db.Queries) ports.UseCasePorts {
	organization := services.NewRepository(db)
	return &UseCase{usecase: organization}
}