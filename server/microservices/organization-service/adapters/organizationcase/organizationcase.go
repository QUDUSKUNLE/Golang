package organizationcase

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/services"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
)

type UseCase struct {
	usecase ports.RepositoryPorts
}

// GetOrganizationID implements ports.OrganizationPorts.
func (u *UseCase) GetOrganizationByUserID(ctx context.Context, user_id string) (*db.Organization, error) {
	return u.usecase.GetOrganizationByUserID(ctx, user_id)
}

// CreateOrganization implements ports.OrganizationPorts.
func (u *UseCase) CreateOrganization(ctx context.Context, user dto.OrganizationDto) (*db.Organization, error) {
	return u.usecase.CreateOrganization(ctx, user)
}

// GetOrganization implements ports.OrganizationPorts.
func (u *UseCase) GetOrganization(ctx context.Context, id string) (*db.Organization, error) {
	return u.usecase.GetOrganization(ctx, id)
}

func InitOrganizationServer(db *db.Queries) ports.OrganizationPorts {
	organization := services.NewRepository(db)
	return &UseCase{usecase: organization}
}
