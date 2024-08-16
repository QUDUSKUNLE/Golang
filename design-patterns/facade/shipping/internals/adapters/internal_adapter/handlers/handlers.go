package handlers

import (
	internalServices "github.com/QUDUSKUNLE/shipping/internals/core/services"
	externalServices "github.com/QUDUSKUNLE/shipping/internals/adapters/external_adapter"
)

type HTTPHandler struct {
	internalServicesAdapter internalServices.InternalServicesHandler
	externalServicesAdapter externalServices.ExternalServicesHandler
}

func HttpAdapter(internalServiceHandler internalServices.InternalServicesHandler, externalServicesHandler externalServices.ExternalServicesHandler) *HTTPHandler {
	return &HTTPHandler{
		internalServicesAdapter: internalServiceHandler,
		externalServicesAdapter: externalServicesHandler,
	}
}
