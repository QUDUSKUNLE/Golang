package services

import (
	"fmt"

	"github.com/QUDUSKUNLE/shipping/internal/core/ledger"
	"github.com/QUDUSKUNLE/shipping/internal/core/model"
	"github.com/QUDUSKUNLE/shipping/internal/core/utils"
)


type LoginAdaptor struct {
	userService *model.User
	userRepositoryService *ledger.UserRepository
	utilsService *utils.Utils
}

func NewLogInAdaptor(loginDto model.LogInDTO) (string, error) {
	fmt.Println("Initiate a new login")
	loginAdaptor := &LoginAdaptor{
		userService: &model.User{},
		userRepositoryService: &ledger.UserRepository{},
		utilsService: &utils.Utils{},
	}
	user, err := loginAdaptor.userRepositoryService.QueryLedgerByEmail(loginDto.Email)
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
