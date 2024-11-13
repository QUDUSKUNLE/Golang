package usecase

import (
	"errors"

	"github.com/QUDUSKUNLE/microservices/services/auth-service/internal/models"
	"gorm.io/gorm"
)

// GetByEmail implements v1.RepoInterface.
func (use *UseCase) GetByEmail(email string) (models.User, error) {
	return use.repo.GetByEmail(email)
}

// Create implements v1.RepoInterface.
func (use *UseCase) Create(user models.User) error {
	return use.repo.Create(user)
}

// Delete implements v1.RepoInterface.
func (use *UseCase) Delete(id string) error {
	return use.repo.Delete(id)
}

// Get implements v1.RepoInterface.
func (use *UseCase) Read(id string) (models.User, error) {
	var (
		user models.User
		err  error
	)
	if user, err = use.repo.Read(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, errors.New("no such user with the id supplied")
		}
		return models.User{}, err
	}
	return user, nil
}

// Update implements v1.RepoInterface.
func (use *UseCase) Update(user models.User) error {
	return use.repo.Update(user)
}
