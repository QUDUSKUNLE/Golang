package shipping

import (
	"time"

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
		return err
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
