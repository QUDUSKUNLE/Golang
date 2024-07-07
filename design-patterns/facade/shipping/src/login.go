package shipping

import (
	// "github.com/QUDUSKUNLE/shipping/src/dto"
	"github.com/QUDUSKUNLE/shipping/src/ledger"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/QUDUSKUNLE/shipping/src/utils"
)


type LoginAdaptor struct {
	userService *model.User
	userRepositoryService *ledger.UserRepository
	utilService *utils.Utils
}

func NewLogInAdaptor() *LoginAdaptor {
	return &LoginAdaptor{
		userService: &model.User{},
		userRepositoryService: &ledger.UserRepository{},
		utilService: &utils.Utils{},
	}
}

func (loginAdaptor *LoginAdaptor) LoginUser(user model.LogInDTO) (string, error) {
	registeredUser, err := loginAdaptor.userRepositoryService.QueryLedgerByEmail(user.Email)
	if err != nil {
		return "", err
	}
	if err := loginAdaptor.userService.ComparePassword(registeredUser.Password, user.Password); err != nil {
		return "", err
	}
	token, err := loginAdaptor.utilService.GenerateAccessToken(registeredUser.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
