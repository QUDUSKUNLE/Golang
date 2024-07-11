package handlers

import (
	services "github.com/QUDUSKUNLE/shipping/internal/core/services"
)

type HTTPHandler struct {
	ServicesAdapter services.ServicesHandler
}

func HttpAdapter(serviceHandlers services.ServicesHandler) *HTTPHandler {
	return &HTTPHandler{
		ServicesAdapter: serviceHandlers,
	}
}
