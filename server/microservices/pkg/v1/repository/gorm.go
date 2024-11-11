package repo

import (
	"github.com/QUDUSKUNLE/microservices/internal/models"
)

// GetByEmail implements v1.RepoInterface.
func (repository *Repository) GetByEmail(email string) (models.User, error) {
	panic("unimplemented")
}

// Create implements v1.RepoInterface.
func (repository *Repository) Create(user models.User) error {
	err := repository.database.Create(&user).Error
	return err
}

// Delete implements v1.RepoInterface.
func (repository *Repository) Delete(id string) error {
	err := repository.database.Where("id = ?", id).Delete(&models.User{}).Error
	return err
}

// Get implements v1.RepoInterface.
func (repository *Repository) Get(id string) (models.User, error) {
	var user models.User
	err := repository.database.Where("id = ?", id).First(&user).Error
	return user, err
}

// Update implements v1.RepoInterface.
func (repository *Repository) Update(user models.User) error {
	panic("unimplemented")
}

// Update implements v1.RepoInterface.
func (repository *Repository) Greet(name string) (string, error) {
	panic("unimplemented")
}
