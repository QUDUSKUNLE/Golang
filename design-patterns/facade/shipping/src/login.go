package shipping

import (
	"github.com/QUDUSKUNLE/shipping/src/dto"
	"github.com/QUDUSKUNLE/shipping/src/ledger"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/QUDUSKUNLE/shipping/src/utils"
)


type LoginAdaptor struct {
	user *model.User
	ledger *ledger.UserLedger
	utils *utils.Utils
}

func NewLogInAdaptor() *LoginAdaptor {
	return &LoginAdaptor{
		ledger: &ledger.UserLedger{},
		user: &model.User{},
		utils: &utils.Utils{},
	}
}

func (loginAdaptor *LoginAdaptor) LoginUser(user dto.LogInDTO) (string, error) {
	registeredUser, err := loginAdaptor.ledger.QueryLedgerByEmail(user.Email)
	if err != nil {
		return "", err
	}
	if err := loginAdaptor.user.ComparePassword(registeredUser.Password, user.Password); err != nil {
		return "", err
	}
	token, err := loginAdaptor.utils.GenerateAccessToken(registeredUser.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
