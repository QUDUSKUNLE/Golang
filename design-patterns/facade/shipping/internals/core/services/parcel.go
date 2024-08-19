package services

import "github.com/QUDUSKUNLE/shipping/internals/core/domain"

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
