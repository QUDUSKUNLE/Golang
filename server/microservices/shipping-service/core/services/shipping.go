package services

import (
	"github.com/QUDUSKUNLE/microservices/shipping-service/core/ports"
)

type ServicesHandler struct {
	internal ports.RepositoryPorts
}

func InternalServicesAdapter(repositoryPort ports.RepositoryPorts) *ServicesHandler {
	return &ServicesHandler{
		internal: repositoryPort,
	}
}
