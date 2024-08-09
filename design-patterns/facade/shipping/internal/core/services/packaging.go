package services

import (
	"fmt"

	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (httpHandler *ExternalServicesHandler) CreatePackagingAdaptor(packaging domain.PackagingDTO) (interface{}, error ){
	fmt.Println("Initiate a new packaging")
	var result interface{}
	result, err := httpHandler.external.CreatePackaging(packaging)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (httpHandler *ExternalServicesHandler) GetPackagingRatesAdaptor() (interface{}, error ){
	fmt.Println("Initiate a new rates")
	var result interface{}
	result, err := httpHandler.external.GetPackagingRates()
	if err != nil {
		return result, err
	}
	return result, nil
}
