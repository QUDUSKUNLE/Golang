package ports

import (
	domain "github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

type ServicePorts interface {
	SaveUserAdaptor(user domain.User) error
	// Log
}

