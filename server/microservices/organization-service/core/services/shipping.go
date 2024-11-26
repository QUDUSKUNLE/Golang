package services

import (
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
)

type ServicesHandler struct {
	internal ports.RepositoryPorts
}

func InternalServicesAdapter(repositoryPort ports.RepositoryPorts) *ServicesHandler {
	return &ServicesHandler{
		internal: repositoryPort,
	}
}
