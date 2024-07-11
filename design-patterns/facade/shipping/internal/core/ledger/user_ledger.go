package ledger

import (
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internal/adapters/repository"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

type Ledger struct {}

func (ledger *Ledger) QueryCreateUser(user *domain.User) error {
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

func (ledger *Ledger) QueryUsers() ([]domain.User, error) {
	// Open repository conection
	db, err := repository.OpenDBConnection()
	if err != nil {
		return []domain.User{}, err
	}
	users, err := db.QueryUsers();
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (ledger *Ledger) QueryUserByID(userID uuid.UUID) (domain.User, error) {
	// Open repository conection
	db, err := repository.OpenDBConnection()
	if err != nil {
		return domain.User{}, err
	}
	user, err := db.QueryUser(userID);
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (ledger *Ledger) QueryLedgerByEmail(email string) (*domain.User, error) {
	// Open repository conection
	db, err := repository.OpenDBConnection()
	if err != nil {
		return &domain.User{}, err
	}
	user, err := db.QueryUserByEmail(email);
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}
