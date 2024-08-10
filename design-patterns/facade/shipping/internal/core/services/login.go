package services

import (
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (internalHandler *InternalServicesHandler) LogInUserAdaptor(loginDto domain.LogInDto) (*domain.User, error) {
	systemsHandler := internalHandler.NewInternalServicesFacade()
	user, err := internalHandler.internal.ReadUserByEmailAdaptor(loginDto.Email)
	if err != nil {
		return &domain.User{}, err
	}
	if err := systemsHandler.userService.ComparePassword(user.Password, loginDto.Password); err != nil {
		return  &domain.User{}, err
	}
	return &domain.User{ID: user.ID, UserType: user.UserType}, nil
}
