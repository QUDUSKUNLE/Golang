package handlers

import (
	services "github.com/QUDUSKUNLE/shipping/internal/core/services"
)

type HTTPHandler struct {
	internalServicesAdapter services.InternalServicesHandler
	externalServicesAdapter services.ExternalServicesHandler
}

func HttpAdapter(serviceHandlers services.InternalServicesHandler, externalServicesHandler services.ExternalServicesHandler) *HTTPHandler {
	return &HTTPHandler{
		internalServicesAdapter: serviceHandlers,
		externalServicesAdapter: externalServicesHandler,
	}
}
