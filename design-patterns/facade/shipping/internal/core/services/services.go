package services

import (
	ports "github.com/QUDUSKUNLE/shipping/internal/core/ports"
)

type ServicesHandler struct {
	Internal ports.RepositoryPorts
}

func ServicesAdapter(repositoryPort ports.RepositoryPorts) *ServicesHandler {
	return &ServicesHandler{
		Internal: repositoryPort,
	}
}
