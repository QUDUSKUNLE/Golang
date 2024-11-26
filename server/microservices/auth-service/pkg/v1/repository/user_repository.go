package repo

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/auth-service/internal/db"
	interfaces "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	database *db.Queries
}

// CreateUser implements v1.UseCaseInterface.
func (u *Repository) CreateUser(ctx context.Context, user db.CreateUserParams) (*db.User, error) {
	return u.database.CreateUser(ctx, user)
}

// GetUser implements v1.UseCaseInterface.
func (u *Repository) GetUser(ctx context.Context, id string) (*db.User, error) {
	return u.database.GetUser(ctx, id)
}

func (u *Repository) GetUserByEmail(ctx context.Context, email pgtype.Text) (*db.User, error) {
	return u.database.GetUserByEmail(ctx, email)
}

func NewRepository(dbase *db.Queries) interfaces.RepositoryInterface {
	return &Repository{database: dbase}
}
