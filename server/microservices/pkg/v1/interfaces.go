package v1

import (
	"github.com/QUDUSKUNLE/microservices/internal/models"
)

type RepositoryInterface interface {
	Create(models.User) error
	Get(id string) (models.User, error)
	Update(models.User) error
	Delete(id string) error
	GetByEmail(email string) (models.User, error)
}


type UseCaseInterface interface {
	Create(user models.User) error
	Get(id string) (models.User, error)
	Update(user models.User) error
	Delete(id string) error
}
