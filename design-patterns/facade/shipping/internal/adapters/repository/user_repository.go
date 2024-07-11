package repository

import (
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internal/core/model"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	*gorm.DB
}

func (database *PostgresRepository) QueryUser(ID uuid.UUID) (model.User, error) {
	user := model.User{ID:  ID}
	result := database.First(&user);
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (database *PostgresRepository) QueryUsers() ([]model.User, error) {
	users := []model.User{}
	if err := database.Find(&users); err != nil {
		return []model.User{}, nil
	}
	return users, nil
}

func (database *PostgresRepository) QueryUserByEmail(email string) (*model.User, error) {
	user := model.User{}
	_ = database.Where(&model.User{Email: email}).First(&user);
	return &user, nil
}

func (database *PostgresRepository) QueryCreateUser(user model.User) error {
	query := model.User{
		Email: user.Email,
		Password: user.Password,
		UserType: user.UserType,
	}
	result := database.Create(&query)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (database *PostgresRepository) QueryUpdateUser(id uuid.UUID, user model.User) error {
	database.Model(&model.User{ID: id}).Updates(model.User{Email: user.Email, Password: user.Password,  UpdatedAt: user.UpdatedAt})
	return nil
}
