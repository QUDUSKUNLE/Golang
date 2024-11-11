package usecase

import (
	"errors"

	"github.com/QUDUSKUNLE/microservices/internal/models"
	"gorm.io/gorm"
)

// GetByEmail implements v1.RepoInterface.
func (use *UseCase) GetByEmail(email string) (models.User, error) {
	panic("Unimplemented")
}

// Create implements v1.RepoInterface.
func (use *UseCase) Create(user models.User) error {
	if _, err := use.repo.GetByEmail(user.Email); !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("the email is already associated with another user")
	}
	return use.repo.Create(user)
}

// Delete implements v1.RepoInterface.
func (use *UseCase) Delete(id string) error {
	panic("Unimplemented")
}

// Get implements v1.RepoInterface.
func (use *UseCase) Get(id string) (models.User, error) {
	var (
		user models.User
		err error
	)
	if user, err = use.repo.Get(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, errors.New("no such user with the id supplied")
		}
		return models.User{}, err
	}
	return user, nil
}

// Update implements v1.RepoInterface.
func (use *UseCase) Update(models.User) error {
	panic("Unimplemented")
}
