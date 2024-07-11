package ledger

import (
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internal/adapters/repository"
	"github.com/QUDUSKUNLE/shipping/internal/core/model"
)

type UserRepository struct {}

func (ledger *UserRepository) QueryCreateUser(user *model.User) error {
	// Open repository conection
	db, err := repository.OpenDBConnection()
	if err != nil {
		return err
	}
	if err := db.QueryCreateUser(*user); err != nil {
		return err
	}
	return nil
}

func (ledger *UserRepository) QueryUsers() ([]model.User, error) {
	// Open repository conection
	db, err := repository.OpenDBConnection()
	if err != nil {
		return []model.User{}, err
	}
	users, err := db.QueryUsers();
	if err != nil {
		return []model.User{}, err
	}
	return users, nil
}

func (ledger *UserRepository) QueryUserByID(userID uuid.UUID) (model.User, error) {
	// Open repository conection
	db, err := repository.OpenDBConnection()
	if err != nil {
		return model.User{}, err
	}
	user, err := db.QueryUser(userID);
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (ledger *UserRepository) QueryLedgerByEmail(email string) (*model.User, error) {
	// Open repository conection
	db, err := repository.OpenDBConnection()
	if err != nil {
		return &model.User{}, err
	}
	user, err := db.QueryUserByEmail(email);
	if err != nil {
		return &model.User{}, err
	}
	return user, nil
}
