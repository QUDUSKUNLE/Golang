package services

import "github.com/QUDUSKUNLE/shipping/internals/core/domain"

type InternalServicesFacade struct {
	userService *domain.User
	notificationService *NotificationService
	locationService *domain.Location
	packagingService *domain.Packaging
	parcelService *domain.Parcel
	labelService *LabelService
	pickUpService *domain.PickUp
	shippingService *domain.Shipping
}

func NewInternalServicesFacade() *InternalServicesFacade {
	return &InternalServicesFacade{
		shippingService: &domain.Shipping{},
		pickUpService: &domain.PickUp{},
		labelService: &LabelService{},
		packagingService: &domain.Packaging{},
		parcelService: &domain.Parcel{},
		userService: &domain.User{},
		notificationService: &NotificationService{},
		locationService: &domain.Location{},
	}
}
