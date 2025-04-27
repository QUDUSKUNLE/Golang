package repo

import (
	"context"

	v1 "github.com/QUDUSKUNLE/microservices/user-service/pkg/v1"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	database *db.Queries
}

// GetUsers implements v1.RepositoryInterface.
func (u *Repository) GetUsers(ctx context.Context) ([]*db.User, error) {
	return u.database.GetUsers(ctx)
}

// CreateUser implements v1.UseCaseInterface.
func (u *Repository) CreateUser(ctx context.Context, user db.CreateUserParams) (*db.CreateUserRow, error) {
	return u.database.CreateUser(ctx, user)
}

// GetUser implements v1.UseCaseInterface.
func (u *Repository) GetUser(ctx context.Context, id string) (*db.User, error) {
	return u.database.GetUser(ctx, id)
}

func (u *Repository) GetUserByEmail(ctx context.Context, email pgtype.Text) (*db.User, error) {
	return u.database.GetUserByEmail(ctx, email)
}

func (u *Repository) UpdateNin(ctx context.Context, user db.UpdateNinParams) (*db.UpdateNinRow, error) {
	return u.database.UpdateNin(ctx, user)
}

func NewRepository(dbase *db.Queries) v1.UserRepository {
	return &Repository{database: dbase}
}
