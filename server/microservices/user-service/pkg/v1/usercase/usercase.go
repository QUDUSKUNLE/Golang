package usercase

import (
	"context"

	v1 "github.com/QUDUSKUNLE/microservices/user-service/pkg/v1"
	repo "github.com/QUDUSKUNLE/microservices/user-service/pkg/v1/repository"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserCase struct {
	repo v1.UserRepository
}

// GetUsers implements v1.UseCaseInterface.
func (u *UserCase) GetUsers(ctx context.Context) ([]*db.User, error) {
	return u.repo.GetUsers(ctx)
}

// UpdateNin implements v1.UseCaseInterface.
func (u *UserCase) UpdateNin(ctx context.Context, data db.UpdateNinParams) (*db.UpdateNinRow, error) {
	return u.repo.UpdateNin(ctx, data)
}

// GetUserByEmail implements v1.RepositoryInterface.
func (u *UserCase) GetUserByEmail(ctx context.Context, email pgtype.Text) (*db.User, error) {
	panic("unimplemented")
}

// Login implements v1.UseCaseInterface.
func (u *UserCase) Login(ctx context.Context, user dto.LogInDto) (*db.User, error) {
	return u.repo.GetUserByEmail(ctx, pgtype.Text{String: user.Email, Valid: true})
}

// CreateUser implements v1.UseCaseInterface.
func (u *UserCase) CreateUser(ctx context.Context, user db.CreateUserParams) (*db.CreateUserRow, error) {
	return u.repo.CreateUser(ctx, user)
}

// GetUser implements v1.UseCaseInterface.
func (u *UserCase) GetUser(ctx context.Context, id string) (*db.User, error) {
	return u.repo.GetUser(ctx, id)
}

func New(repo v1.UserRepository) v1.UserRepository {
	return &UserCase{repo: repo}
}

func InitUserServer(db *db.Queries) v1.UserPorts {
	userRepo := repo.NewRepository(db)
	return &UserCase{repo: userRepo}
}
