package repo

import (
	"github.com/QUDUSKUNLE/microservices/auth-service/internal/models"
)

// Create implements v1.RepoInterface.
func (repository *Repository) CreateUser(user models.User) error {
	return repository.database.Create(&user).Error
}

// Get implements v1.RepoInterface.
func (repository *Repository) GetUsers() ([]*models.User, error) {
	var user []*models.User
	err := repository.database.Limit(50).Find(&user).Error
	return user, err
}

// Get implements v1.RepoInterface.
func (repository *Repository) GetUser(id string) (models.User, error) {
	var user models.User
	err := repository.database.Where("id = ?", id).First(&user).Error
	return user, err
}

// Get implements v1.RepoInterface.
func (repository *Repository) LogIn(user models.LogInDto) (models.User, error) {
	return repository.GetByEmail(user.Email)
}

// Get implements v1.RepoInterface.
func (repository *Repository) Reads() ([]models.User, error) {
	var user []models.User
	err := repository.database.Limit(10).Find(&user).Error
	return user, err
}

// GetByEmail implements v1.RepoInterface.
func (repository *Repository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := repository.database.Where("email = ?", email).Find(&user).Error
	return user, err
}

// Delete implements v1.RepoInterface.
func (repository *Repository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&models.User{}).Error
}

// Update implements v1.RepoInterface.
func (repository *Repository) Update(user models.User) error {
	return repository.database.Where("id = ?", user.ID).Updates(&user).Error
}
