package services

import (
	"fmt"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (httpHandler *InternalServicesHandler) LogInUserAdaptor(loginDto domain.LogInDTO) (*domain.User, error) {
	fmt.Println("Initiate a new login")
	systemsHandler := httpHandler.NewInternalServicesFacade()
	user, err := httpHandler.internal.ReadUserByEmailAdaptor(loginDto.Email)
	if err != nil {
		return &domain.User{}, err
	}
	if err := systemsHandler.userService.ComparePassword(user.Password, loginDto.Password); err != nil {
		return  &domain.User{}, err
	}
	return &domain.User{ID: user.ID, UserType: user.UserType}, nil
}
