package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (httpHandler *InternalServicesHandler) CarrierPickUpsAdaptor(ID uuid.UUID) ([]domain.PickUp, error) {
	fmt.Println("Get Carrier pick ups")
	pickUps, err := httpHandler.internal.CarrierPickUps(ID)
	if err != nil {
		return []domain.PickUp{}, err
	}
	return pickUps, nil
}
