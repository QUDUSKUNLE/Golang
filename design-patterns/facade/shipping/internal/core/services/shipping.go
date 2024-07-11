package services

import (
	"time"
	"errors"
	"github.com/QUDUSKUNLE/shipping/internal/core/ledger"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/QUDUSKUNLE/shipping/internal/core/utils"
	"github.com/labstack/echo/v4"
	"fmt"
)

type ShippingAdaptor struct {
	utilsService *utils.Utils
	shippingService *domain.Shipping
	shippingRepositoryService *ledger.Ledger
}

func NewShippingAdaptor(cont echo.Context, shippingDto *domain.ShippingDTO) error {
	fmt.Println("Initiate a new shipping")
	adaptor := &ShippingAdaptor{
		shippingService: &domain.Shipping{},
		shippingRepositoryService: &ledger.Ledger{},
		utilsService: &utils.Utils{},
	}
	// Validate user
	user, err := adaptor.utilsService.ParseUserID(cont)
	if err != nil {
		return err
	}

	if user.UserType != string(domain.USER) {
		return errors.New("unauthorized to perform this operation")
	}
	newShipping := adaptor.shippingService.BuildNewShipping(user.ID, *shippingDto)
	err = adaptor.shippingRepositoryService.ShippingLedger(*newShipping)
	if err != nil {
		return err
	}
	pickUpDTO := domain.PickUpDTO{
		ShippingID: newShipping.ID,
		CarrierID: newShipping.UserID,
		Status: string(domain.SCHEDULED),
		PickUpAt: time.Now(),
	}
	err = NewPickUpAdaptor(pickUpDTO)
	if err != nil {
		return err
	}
	fmt.Println("Shipping created successfully.")
	return nil
}

func GetShippingsAdaptor(context echo.Context) ([]domain.Shipping, error) {
	fmt.Println("Initiate a shipping")
	adaptor := &ShippingAdaptor{
		utilsService: &utils.Utils{},
		shippingRepositoryService: &ledger.Ledger{},
	}
	user, err := adaptor.utilsService.ParseUserID(context)
	if err != nil {
		return []domain.Shipping{}, err
	}
	var status string
	status = context.QueryParams().Get("status")
	if status == "" {
		status = "SCHEDULED"
	}
	shippings, err := adaptor.shippingRepositoryService.QueryShippingLedger(user.ID, status)
	if err != nil {
		return []domain.Shipping{}, err
	}
	return shippings, nil
}
