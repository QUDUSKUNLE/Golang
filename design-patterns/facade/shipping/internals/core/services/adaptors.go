package services

import "github.com/QUDUSKUNLE/shipping/internals/core/domain"

type InternalServicesFacade struct {
	userService *domain.User
	notificationService *Notification
	locationService *domain.Location
	packagingService *domain.Packaging
	parcelService *domain.Parcel
	labelService *LabelService
	pickUpService *domain.PickUp
	shippingService *domain.Shipping
}

func (internalServicesHandler *InternalServicesHandler) NewInternalServicesFacade() *InternalServicesFacade {
	return &InternalServicesFacade{
		shippingService: &domain.Shipping{},
		pickUpService: &domain.PickUp{},
		labelService: &LabelService{},
		packagingService: &domain.Packaging{},
		parcelService: &domain.Parcel{},
		userService: &domain.User{},
		notificationService: &Notification{},
		locationService: &domain.Location{},
	}
}
