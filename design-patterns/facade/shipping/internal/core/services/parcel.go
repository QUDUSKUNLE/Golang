package services

import "github.com/QUDUSKUNLE/shipping/internal/core/domain"

func (internalHandler *InternalServicesHandler) NewParcelAdaptor(packageDto domain.ParcelDto) error {
	systemsHandler := internalHandler.NewInternalServicesFacade()
	parcel := systemsHandler.parcelService.BuildNewParcel(packageDto)
	err := internalHandler.internal.SaveParcelAdaptor(*parcel);
	if err != nil {
		return err
	}
	systemsHandler.notificationService.SendAddressNotification()
	return nil
}
