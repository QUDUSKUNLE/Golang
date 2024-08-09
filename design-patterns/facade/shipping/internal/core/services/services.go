package services

import (
	ports "github.com/QUDUSKUNLE/shipping/internal/core/ports"
)

type InternalServicesHandler struct {
	internal ports.RepositoryPorts
}

type ExternalServicesHandler struct {
	external ports.ExternalPorts
}

func InternalServicesAdapter(repositoryPort ports.RepositoryPorts) *InternalServicesHandler {
	return &InternalServicesHandler{
		internal: repositoryPort,
	}
}

func ExternalServicesAdapter(externalPorts ports.ExternalPorts) *ExternalServicesHandler {
	return &ExternalServicesHandler{
		external: externalPorts,
	}
}
