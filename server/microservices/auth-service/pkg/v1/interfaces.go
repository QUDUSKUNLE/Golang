package v1

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/auth-service/internal/db"
	"github.com/QUDUSKUNLE/microservices/auth-service/internal/dto"
	"github.com/jackc/pgx/v5/pgtype"
)

type RepositoryInterface interface {
	CreateUser(ctx context.Context, user db.CreateUserParams) (*db.User, error)
	GetUser(ctx context.Context, id string) (*db.User, error)
	GetUserByEmail(ctx context.Context, email pgtype.Text) (*db.User, error)
	UpdateNin(ctx context.Context, user db.UpdateNinParams) (*db.User, error)
}

type UseCaseInterface interface {
	CreateUser(ctx context.Context, user db.CreateUserParams) (*db.User, error)
	GetUser(ctx context.Context, id string) (*db.User, error)
	Login(ctx context.Context, user dto.LogInDto) (*db.User, error)
	UpdateNin(ctx context.Context, user db.UpdateNinParams) (*db.User, error)
}
