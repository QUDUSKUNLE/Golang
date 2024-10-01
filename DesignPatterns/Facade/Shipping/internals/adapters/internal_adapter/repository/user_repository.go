package repository

import (
	"errors"

	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (database *PostgresRepository) ReadUserAdaptor(ID uuid.UUID) (*domain.User, error) {
	user := domain.User{ID: ID}
	err := database.db.First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.User{}, err
	}
	return &user, nil
}

func (database *PostgresRepository) QueryUsers() ([]*domain.User, error) {
	users := []*domain.User{}
	if err := database.db.Find(&users).Limit(10); err != nil {
		return []*domain.User{}, nil
	}
	return users, nil
}

func (database *PostgresRepository) ReadUserByEmailAdaptor(email string) (*domain.User, error) {
	user := domain.User{}
	result := database.db.Where(&domain.User{Email: email}).First(&user).Limit(1)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
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
