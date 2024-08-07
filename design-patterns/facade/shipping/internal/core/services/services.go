package services

import (
	ports "github.com/QUDUSKUNLE/shipping/internal/core/ports"
)

type ServicesHandler struct {
	internal ports.RepositoryPorts
}

func ServicesAdapter(repositoryPort ports.RepositoryPorts) *ServicesHandler {
	return &ServicesHandler{
		internal: repositoryPort,
	}
}
