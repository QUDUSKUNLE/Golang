package handlers

import (
	services "github.com/QUDUSKUNLE/shipping/internal/core/services"
)

type HTTPHandler struct {
	servicesAdapter services.ServicesHandler
}

func HttpAdapter(serviceHandlers services.ServicesHandler) *HTTPHandler {
	return &HTTPHandler{
		servicesAdapter: serviceHandlers,
	}
}
