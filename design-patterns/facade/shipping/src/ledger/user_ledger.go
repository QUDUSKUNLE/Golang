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
