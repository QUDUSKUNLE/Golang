package ledger

import (
	"errors"

	"github.com/QUDUSKUNLE/shipping/src/database"
	"github.com/QUDUSKUNLE/shipping/src/model"
)

type UserRepository struct {}

func (ledger *UserRepository) QueryCreateUser(user *model.User) error {
	// Open database conection
	db, err := database.OpenDBConnection()
	if err != nil {
		return err
	}
	if err := db.QueryCreateUser(*user); err != nil {
		return errors.New("user`s already exist")
	}
	return nil
}

func (ledger *UserRepository) QueryLedger(id string) (model.User, error) {
	// Open database conection
	db, err := database.OpenDBConnection()
	if err != nil {
		return model.User{}, err
	}
	user, err := db.QueryUser(id);
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (ledger *UserRepository) QueryLedgerByEmail(email string) (model.User, error) {
	// Open database conection
	db, err := database.OpenDBConnection()
	if err != nil {
		return model.User{}, err
	}
	user, err := db.QueryUserByEmail(email);
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
