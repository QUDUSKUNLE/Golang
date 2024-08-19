package services

import (
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
)

func (internalHandler *InternalServicesHandler) NewPackagingAdaptor(packageDto domain.PackagingDto) error {
	systemsHandler := NewInternalServicesFacade()
	pack := systemsHandler.packagingService.BuildNewPackaging(packageDto)
	err := internalHandler.internal.SavePackagingAdaptor(pack);
	if err != nil {
		return err
	}
	systemsHandler.notificationService.SendPackagingNotification()
	return nil
}
