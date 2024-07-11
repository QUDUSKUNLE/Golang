package services

import (
	"time"
	"errors"
	"github.com/QUDUSKUNLE/shipping/internal/core/ledger"
	"github.com/QUDUSKUNLE/shipping/internal/core/model"
	"github.com/QUDUSKUNLE/shipping/internal/core/utils"
	"github.com/labstack/echo/v4"
	"fmt"
)

type ShippingAdaptor struct {
	utilsService *utils.Utils
	shippingService *model.Shipping
	shippingRepositoryService *ledger.ShippingRepository
}

func NewShippingAdaptor(cont echo.Context, shippingDto *model.ShippingDTO) error {
	fmt.Println("Initiate a new shipping")
	adaptor := &ShippingAdaptor{
		shippingService: &model.Shipping{},
		shippingRepositoryService: &ledger.ShippingRepository{},
		utilsService: &utils.Utils{},
	}
	// Validate user
	user, err := adaptor.utilsService.ParseUserID(cont)
	if err != nil {
		return err
	}

	if user.UserType != string(model.USER) {
		return errors.New("unauthorized to perform this operation")
	}
	newShipping := adaptor.shippingService.BuildNewShipping(user.ID, *shippingDto)
	err = adaptor.shippingRepositoryService.ShippingLedger(*newShipping)
	if err != nil {
		return err
	}
	pickUpDTO := model.PickUpDTO{
		ShippingID: newShipping.ID,
		CarrierID: newShipping.UserID,
		Status: string(model.SCHEDULED),
		PickUpAt: time.Now(),
	}
	err = NewPickUpAdaptor(pickUpDTO)
	if err != nil {
		return err
	}
	fmt.Println("Shipping created successfully.")
	return nil
}

func GetShippingsAdaptor(context echo.Context) ([]model.Shipping, error) {
	fmt.Println("Initiate a shipping")
	adaptor := &ShippingAdaptor{
		utilsService: &utils.Utils{},
		shippingRepositoryService: &ledger.ShippingRepository{},
	}
	user, err := adaptor.utilsService.ParseUserID(context)
	if err != nil {
		return []model.Shipping{}, err
	}
	var status string
	status = context.QueryParams().Get("status")
	if status == "" {
		status = "SCHEDULED"
	}
	shippings, err := adaptor.shippingRepositoryService.QueryShippingLedger(user.ID, status)
	if err != nil {
		return []model.Shipping{}, err
	}
	return shippings, nil
}
