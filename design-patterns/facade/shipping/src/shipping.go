package shipping

import (
	"time"
	"errors"
	"github.com/QUDUSKUNLE/shipping/src/ledger"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/QUDUSKUNLE/shipping/src/utils"
	"github.com/google/uuid"
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
	userID, err := uuid.Parse(adaptor.utilsService.ObtainUser(cont))
	if err != nil {
		fmt.Println(err, "&&&")
		return err
	}

	if userID.String() == "" {
		return errors.New("unauthorized")
	}
	newShipping := adaptor.shippingService.BuildNewShipping(userID, *shippingDto)
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
	userID, err := uuid.Parse(adaptor.utilsService.ObtainUser(context))
	if err != nil {
		return []model.Shipping{}, err
	}
	var status string
	status = context.QueryParams().Get("status")
	if status == "" {
		status = "SCHEDULED"
	}
	shippings, err := adaptor.shippingRepositoryService.QueryShippingLedger(userID, status)
	if err != nil {
		return []model.Shipping{}, err
	}
	return shippings, nil
}
