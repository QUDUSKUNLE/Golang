package services

import (
	ports "github.com/QUDUSKUNLE/shipping/internal/core/ports"
)

type ServicesHandler struct {
	Internal ports.ServicePorts
}

func ServicesAdapter(servicePorts ports.ServicePorts) *ServicesHandler {
	return &ServicesHandler{
		Internal: servicePorts,
	}
}
