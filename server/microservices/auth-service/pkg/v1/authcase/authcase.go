package authcase

import (
	"context"

	v1 "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1"
	repo "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1/authrepository"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthCase struct {
	repo v1.AuthRepository
}


// GetUserByEmail implements v1.RepositoryInterface.
func (u *AuthCase) GetUserByEmail(ctx context.Context, email pgtype.Text) (*db.User, error) {
	panic("unimplemented")
}

// Login implements v1.UserPorts.
func (u *AuthCase) Login(ctx context.Context, user dto.LogInDto) (*db.User, error) {
	return u.repo.GetUserByEmail(ctx, pgtype.Text{String: user.Email, Valid: true})
}

// GetUser implements v1.UserPorts.
func (u *AuthCase) GetUser(ctx context.Context, id string) (*db.User, error) {
	return u.repo.GetUser(ctx, id)
}

func New(repo v1.AuthRepository) v1.AuthRepository {
	return &AuthCase{repo: repo}
}

func InitAuthServer(db *db.Queries) v1.AuthPorts {
	userRepo := repo.NewRepository(db)
	return &AuthCase{repo: userRepo}
}
