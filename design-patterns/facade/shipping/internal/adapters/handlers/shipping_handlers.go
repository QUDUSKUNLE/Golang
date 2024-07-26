package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/labstack/echo/v4"
)

func (handler *HTTPHandler) NewShipping(context echo.Context) error {
	shippingDto := new(domain.ShippingDTO)
	if err := handler.ValidateStruct(context, shippingDto); err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}

	// Validate carrier
	user, err := handler.ParseUserID(context)
	if err != nil {
		return context.JSON(http.StatusUnauthorized, echo.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}

	if user.UserType != string(domain.USER) {
		return context.JSON(http.StatusUnauthorized, echo.Map{
			"Message": "Unauthorized to perform this operation",
			"Success": false,
		})
	}

	shippingDto.UserID = user.ID
	// Initiate a new shipping
	err = handler.ServicesAdapter.NewShippingAdaptor(shippingDto);
	if err != nil {
		return context.JSON(http.StatusNotAcceptable, echo.Map{"Message": err.Error(), "Success": false })
	}
	return context.JSON(http.StatusOK, echo.Map{
		"Message": "Product is scheduled for shipping.",
		"Success": true,
	})
}

func (handler *HTTPHandler) GetShippings(context echo.Context) error {
	user, err := handler.ParseUserID(context)
	if err != nil {
		return context.JSON(http.StatusUnauthorized, echo.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}

	if user.UserType != string(domain.USER) {
		return context.JSON(http.StatusUnauthorized, echo.Map{
			"Message": "Unauthorized to perform this operation.",
			"Success": false,
		})
	}

	shippings, err := handler.ServicesAdapter.GetShippingsAdaptor(user.ID);
	if err != nil {
		return context.JSON(http.StatusNotImplemented, echo.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}
	return context.JSON(http.StatusOK, echo.Map{
		"Shippings": shippings,
		"Success": true,
	})
}

func (handler *HTTPHandler) RejectProduct(context echo.Context) error {
	return context.JSON(http.StatusOK, echo.Map{
		"Message": "Product is delivered.",
		"Success": true,
	})
}
