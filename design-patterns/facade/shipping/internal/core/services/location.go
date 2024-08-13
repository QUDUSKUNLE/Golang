package services

import (
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (internalHandler *InternalServicesHandler) NewLocationAdaptor(locationDto domain.LocationDto) error {
	systemsHandler := internalHandler.NewInternalServicesFacade()
	locations := systemsHandler.locationService.BuildNewLocation(locationDto)
	err := internalHandler.internal.SaveAddressAdaptor(locations);
	if err != nil {
		return err
	}
	systemsHandler.notificationService.SendAddressNotification()
	return nil
}

func (internalHandler *InternalServicesHandler) GetLocationAdaptor(addressID, userID uuid.UUID) (*domain.Location, error) {
	location, err := internalHandler.internal.ReadAddressAdaptor(addressID, userID);
	if err != nil {
		return &domain.Location{}, err
	}
	return location, nil
}

func (internalHandler *InternalServicesHandler) GetLocationsAdaptor(userID uuid.UUID) ([]domain.Location, error) {
	location, err := internalHandler.internal.ReadAddressesAdaptor(userID);
	if err != nil {
		return []domain.Location{}, err
	}
	return location, nil
}

func (internalHandler *InternalServicesHandler) TerminalUpdateAddressAdaptor(location domain.Location) error {
	err := internalHandler.internal.TerminalUpdateAddressAdaptor(location);
	if err != nil {
		return err
	}
	return nil
}
