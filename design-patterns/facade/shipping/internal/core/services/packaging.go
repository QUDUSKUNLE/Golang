package services

import (
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (internalHandler *InternalServicesHandler) NewPackagingAdaptor(packageDto domain.PackagingDto) error {
	systemsHandler := internalHandler.NewInternalServicesFacade()
	pack := systemsHandler.packagingService.BuildNewPackaging(packageDto)
	err := internalHandler.internal.SavePackagingAdaptor(*pack);
	if err != nil {
		return err
	}
	systemsHandler.notificationService.SendAddressNotification()
	return nil
}
