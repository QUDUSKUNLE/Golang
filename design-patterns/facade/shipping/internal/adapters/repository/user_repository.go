package repository

import (
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/google/uuid"
)


func (database *PostgresRepository) ReadUserAdaptor(ID uuid.UUID) (*domain.User, error) {
	user := domain.User{ID: ID}
	result := database.db.First(&user)
	if result.Error != nil {
		return &domain.User{}, result.Error
	}
	return &user, nil
}

func (database *PostgresRepository) QueryUsers() ([]domain.User, error) {
	users := []domain.User{}
	if err := database.db.Find(&users); err != nil {
		return []domain.User{}, nil
	}
	return users, nil
}

func (database *PostgresRepository) ReadUserByEmailAdaptor(email string) (*domain.User, error) {
	user := domain.User{}
	result := database.db.Where(&domain.User{Email: email}).First(&user)
	if result.Error != nil {
		return &user, result.Error
	}
	return &user, nil
}

func (database *PostgresRepository) SaveUserAdaptor(user domain.User) error {
	_ = database.db.AutoMigrate(&domain.User{})
	query := domain.User{
		Email:    user.Email,
		Password: user.Password,
		UserType: user.UserType,
	}
	result := database.db.Create(&query)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (database *PostgresRepository) QueryUpdateUser(id uuid.UUID, user domain.User) error {
	database.db.Model(&domain.User{ID: id}).Updates(domain.User{Email: user.Email, Password: user.Password, UpdatedAt: user.UpdatedAt})
	return nil
}
