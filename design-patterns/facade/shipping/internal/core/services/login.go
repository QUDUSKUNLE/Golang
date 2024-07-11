package services

import (
	"fmt"

	"github.com/QUDUSKUNLE/shipping/internal/core/ledger"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/QUDUSKUNLE/shipping/internal/core/utils"
)


type LoginAdaptor struct {
	userService *domain.User
	ledger *ledger.Ledger
	utilsService *utils.Utils
}

func (httpHandler *ServicesHandler) LogInUserAdaptor(loginDto domain.LogInDTO) (string, error) {
	fmt.Println("Initiate a new login")
	loginAdaptor := &LoginAdaptor{
		userService: &domain.User{},
		ledger: &ledger.Ledger{},
		utilsService: &utils.Utils{},
	}
	user, err := loginAdaptor.ledger.QueryLedgerByEmail(loginDto.Email)
	if err != nil {
		return "", err
	}
	if err := loginAdaptor.userService.ComparePassword(user.Password, loginDto.Password); err != nil {
		return "", err
	}
	token, err := loginAdaptor.utilsService.GenerateAccessToken(*user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewLogInAdaptor(loginDto domain.LogInDTO) (string, error) {
	fmt.Println("Initiate a new login")
	loginAdaptor := &LoginAdaptor{
		userService: &domain.User{},
		ledger: &ledger.Ledger{},
		utilsService: &utils.Utils{},
	}
	user, err := loginAdaptor.ledger.QueryLedgerByEmail(loginDto.Email)
	if err != nil {
		return "", err
	}
	if err := loginAdaptor.userService.ComparePassword(user.Password, loginDto.Password); err != nil {
		return "", err
	}
	token, err := loginAdaptor.utilsService.GenerateAccessToken(*user)
	if err != nil {
		return "", err
	}
	return token, nil
}
