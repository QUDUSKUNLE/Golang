package v1

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user db.CreateUserParams) (*db.CreateUserRow, error)
	GetUser(ctx context.Context, id string) (*db.User, error)
	GetUsers(ctx context.Context) ([]*db.User, error)
	GetUserByEmail(ctx context.Context, email pgtype.Text) (*db.User, error)
	UpdateNin(ctx context.Context, user db.UpdateNinParams) (*db.UpdateNinRow, error)
}

type UserPorts interface {
	CreateUser(ctx context.Context, user db.CreateUserParams) (*db.CreateUserRow, error)
	GetUser(ctx context.Context, id string) (*db.User, error)
	GetUsers(ctx context.Context) ([]*db.User, error)
	Login(ctx context.Context, user dto.LogInDto) (*db.User, error)
	UpdateNin(ctx context.Context, user db.UpdateNinParams) (*db.UpdateNinRow, error)
}
