package services

import (
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (internalHandler *InternalServicesHandler) CarrierPickUpsAdaptor(ID uuid.UUID) ([]domain.PickUp, error) {
	pickUps, err := internalHandler.internal.CarrierPickUps(ID)
	if err != nil {
		return []domain.PickUp{}, err
	}
	return pickUps, nil
}
