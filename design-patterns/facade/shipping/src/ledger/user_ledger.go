package ledger

import (
	"errors"

	"github.com/QUDUSKUNLE/shipping/src/database"
	"github.com/QUDUSKUNLE/shipping/src/model"
)

type UserLedger struct {}

func (ledger *UserLedger) UserLedger(user *model.User) error {
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

func (ledger *UserLedger) QueryLedger(id string) (model.User, error) {
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

func (ledger *UserLedger) QueryLedgerByEmail(email string) (model.User, error) {
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
