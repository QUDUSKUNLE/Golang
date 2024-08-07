package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

type LocationAdaptor struct {
	locationService *domain.Location
	notificationService *Notification
}

func (httpHandler *ServicesHandler) NewLocationAdaptor(locationDto []domain.LocationDTO) error {
	fmt.Println("Initiate new addresses savings")
	adaptor := &LocationAdaptor{
		locationService: &domain.Location{},
		notificationService: &Notification{},
	}
	locations := adaptor.locationService.BuildNewLocation(locationDto)
	err := httpHandler.Internal.SaveAddressAdaptor(locations);
	if err != nil {
		return err
	}
	adaptor.notificationService.SendAddressNotification()
	fmt.Println("New addresses saved successfully.")
	return nil
}

func (httpHandler *ServicesHandler) GetLocationAdaptor(addressID, userID uuid.UUID) (*domain.Location, error) {
	fmt.Println("Get a address")
	location, err := httpHandler.Internal.ReadAddressAdaptor(addressID, userID);
	if err != nil {
		return &domain.Location{}, err
	}
	return location, nil
}

func (httpHandler *ServicesHandler) GetLocationsAdaptor(userID uuid.UUID) ([]domain.Location, error) {
	fmt.Println("Get addresses")
	location, err := httpHandler.Internal.ReadAddressesAdaptor(userID);
	if err != nil {
		return []domain.Location{}, err
	}
	return location, nil
}
