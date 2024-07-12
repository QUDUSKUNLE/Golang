package services

import (
	"fmt"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)


type LoginAdaptor struct {
	userService *domain.User
}

func (httpHandler *ServicesHandler) LogInUserAdaptor(loginDto domain.LogInDTO) (*domain.User, error) {
	fmt.Println("Initiate a new login")
	loginAdaptor := &LoginAdaptor{
		userService: &domain.User{},
	}
	user, err := httpHandler.Internal.ReadUserByEmail(loginDto.Email)
	if err != nil {
		return &domain.User{}, err
	}
	if err := loginAdaptor.userService.ComparePassword(user.Password, loginDto.Password); err != nil {
		return  &domain.User{}, err
	}
	return &domain.User{ID: user.ID, UserType: user.UserType}, nil
}
