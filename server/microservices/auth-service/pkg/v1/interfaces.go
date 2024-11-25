package v1

import (
	"github.com/QUDUSKUNLE/microservices/auth-service/internal/models"
)

type RepositoryInterface interface {
	CreateUser(models.User) error
	GetUsers() ([]*models.User, error)
	GetUser(id string) (models.User, error)
	Update(models.User) error
	Delete(id string) error
	GetByEmail(email string) (models.User, error)
	LogIn(user models.LogInDto) (models.User, error)
}

type UseCaseInterface interface {
	CreateUser(user models.User) error
	GetUsers() ([]*models.User, error)
	GetUser(id string) (models.User, error)
	Update(user models.User) error
	Delete(id string) error
	LogIn(user models.LogInDto) (models.User, error)
}
