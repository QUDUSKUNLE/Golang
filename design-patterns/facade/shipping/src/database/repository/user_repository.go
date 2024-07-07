package repository

import (
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	*sqlx.DB
}

func (database *Database) QueryUser(ID string) (model.User, error) {
	user := model.User{}
	query := `SELECT * FROM users WHERE id=$1`
	if err := database.Get(&user, query, ID); err != nil {
		return model.User{}, nil
	}
	return user, nil
}

func (database *Database) QueryUserByEmail(email string) (model.User, error) {
	user := model.User{}
	query := `SELECT * FROM public.users WHERE email=$1`
	if err := database.Get(&user, query, email); err != nil {
		return model.User{}, nil
	}
	return user, nil
}

func (database *Database) QueryCreateUser(user model.User) error {
	query := `INSERT INTO users (email, pass, user_type) VALUES ($1, $2, $3)`
	_, err := database.Exec(query, user.Email, user.Password, user.UserType)
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) QueryUpdateUser(id string, user model.User) error {
	q := `UPDATE user SET email=$2, pass=$3, created_at=$4, updated_at=$5 WHERE id=$1`;
	_, err := database.Exec(q, id, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
