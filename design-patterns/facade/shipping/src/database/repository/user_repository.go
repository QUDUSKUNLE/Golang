package repository

import (
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func (database *Database) QueryUser(ID uuid.UUID) (model.User, error) {
	user := model.User{ID:  ID}
	if err := database.First(&user); err != nil {
		return model.User{}, nil
	}
	return user, nil
}

func (database *Database) QueryUserByEmail(email string) (model.User, error) {
	user := model.User{Email: email}
	if err := database.First(&user); err != nil {
		return model.User{}, nil
	}
	return user, nil
}

func (database *Database) QueryCreateUser(user model.User) error {
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

func (database *Database) QueryUpdateUser(id uuid.UUID, user model.User) error {
	database.Model(&model.User{ID: id}).Updates(model.User{Email: user.Email, Password: user.Password,  UpdatedAt: user.UpdatedAt})
	return nil
}
