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
func (use *UseCase) CreateUser(user models.User) error {
	return use.repo.CreateUser(user)
}

// Create implements v1.RepoInterface.
func (use *UseCase) GetUsers() ([]*models.User, error) {
	return use.repo.GetUsers()
}

// Create implements v1.RepoInterface.
func (use *UseCase) LogIn(user models.LogInDto) (models.User, error) {
  registeredUser, err := use.repo.LogIn(user)
	if err != nil {
		return models.User{}, err
	}
	if err = registeredUser.ComparePassword(user.Password); err != nil {
		return models.User{}, err
	}
	return registeredUser, nil
}

// Delete implements v1.RepoInterface.
func (use *UseCase) Delete(id string) error {
	return use.repo.Delete(id)
}

// Get implements v1.RepoInterface.
func (use *UseCase) GetUser(id string) (models.User, error) {
	var (
		user models.User
		err  error
	)
	if user, err = use.repo.GetUser(id); err != nil {
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
