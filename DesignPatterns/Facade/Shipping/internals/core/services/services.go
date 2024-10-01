package services

import (
	ports "github.com/QUDUSKUNLE/shipping/internals/core/ports"
)

type InternalServicesHandler struct {
	internal ports.RepositoryPorts
}

func InternalServicesAdapter(repositoryPort ports.RepositoryPorts) *InternalServicesHandler {
	return &InternalServicesHandler{
		internal: repositoryPort,
	}
}
