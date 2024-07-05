package ledger

import (
	"fmt"

	// "github.com/QUDUSKUNLE/shipping/src/database"
	"github.com/QUDUSKUNLE/shipping/src/model"
)

type UserLedger struct {}

func (ledger *UserLedger) RegisterLedger(user *model.User) error {
	// Open database conection
	// db, err := database.OpenDBConnection()
	// if err != nil {
	// 	return err
	// }
	// if err := db.QueryCreateUser(*user); err != nil {
	// 	return err
	// }
	fmt.Println(user)
	fmt.Printf("Register a new user\nemail:%s\npassword:%s\nid:%s\n", user.Email, user.Password, user.ID)
	return nil
}
