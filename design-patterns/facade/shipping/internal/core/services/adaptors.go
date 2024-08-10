package services

import "github.com/QUDUSKUNLE/shipping/internal/core/domain"

type InternalServicesFacade struct {
	userService *domain.User
	notificationService *Notification
	locationService *domain.Location
	labelService *LabelService
	pickUpService *domain.PickUp
	shippingService *domain.Shipping
}

type ExternalServicesFacade struct {
	terminalService *domain.PackagingDto
}

func (internalServicesHandler *InternalServicesHandler) NewInternalServicesFacade() *InternalServicesFacade {
	return &InternalServicesFacade{
		shippingService: &domain.Shipping{},
		pickUpService: &domain.PickUp{},
		labelService: &LabelService{},
		userService: &domain.User{},
		notificationService: &Notification{},
		locationService: &domain.Location{},
	}
}

func (externalServicesHandler *ExternalServicesHandler) NewExternalServicesFacade() *ExternalServicesFacade {
	return &ExternalServicesFacade{
		terminalService: &domain.PackagingDto{},
	}
}
