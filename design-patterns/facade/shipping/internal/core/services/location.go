package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (httpHandler *InternalServicesHandler) NewLocationAdaptor(locationDto domain.LocationDTO) error {
	fmt.Println("Initiate new addresses savings")
	systemsHandler := httpHandler.NewInternalServicesFacade()
	locations := systemsHandler.locationService.BuildNewLocation(locationDto)
	err := httpHandler.internal.SaveAddressAdaptor(locations);
	if err != nil {
		return err
	}
	systemsHandler.notificationService.SendAddressNotification()
	fmt.Println("New addresses saved successfully.")
	return nil
}

func (httpHandler *InternalServicesHandler) GetLocationAdaptor(addressID, userID uuid.UUID) (*domain.Location, error) {
	fmt.Println("Get a address")
	location, err := httpHandler.internal.ReadAddressAdaptor(addressID, userID);
	if err != nil {
		return &domain.Location{}, err
	}
	return location, nil
}

func (httpHandler *InternalServicesHandler) GetLocationsAdaptor(userID uuid.UUID) ([]domain.Location, error) {
	fmt.Println("Get addresses")
	location, err := httpHandler.internal.ReadAddressesAdaptor(userID);
	if err != nil {
		return []domain.Location{}, err
	}
	return location, nil
}
