package services

import (
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
)

func (internalHandler *InternalServicesHandler) NewParcelAdaptor(packageDto domain.ParcelDto) error {
	systemsHandler := NewInternalServicesFacade()
	parcel := systemsHandler.parcelService.BuildNewParcel(packageDto)
	err := internalHandler.internal.SaveParcelAdaptor(parcel);
	if err != nil {
		return err
	}
	systemsHandler.notificationService.SendParcelNotification()
	return nil
}

func (internalHandler *InternalServicesHandler) GetParcelsAdaptor(userID uuid.UUID) ([]*domain.Parcel, error) {
	systemsHandler := NewInternalServicesFacade()
	parcels, err := internalHandler.internal.GetParcelsAdaptor(userID)
	if err != nil {
		return []*domain.Parcel{}, err
	}
	systemsHandler.notificationService.SendParcelNotification()
	return parcels, nil
}
