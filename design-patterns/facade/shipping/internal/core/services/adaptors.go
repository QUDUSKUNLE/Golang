package services

import "github.com/QUDUSKUNLE/shipping/internal/core/domain"

type ServicesFacade struct {
	userService *domain.User
	notificationService *Notification
	locationService *domain.Location
	labelService *LabelService
	pickUpService *domain.PickUp
	shippingService *domain.Shipping
}

func (httpHandler *ServicesHandler) NewServicesFacade() *ServicesFacade {
	return &ServicesFacade{
		shippingService: &domain.Shipping{},
		pickUpService: &domain.PickUp{},
		labelService: &LabelService{},
		userService: &domain.User{},
		notificationService: &Notification{},
		locationService: &domain.Location{},
	}
}
