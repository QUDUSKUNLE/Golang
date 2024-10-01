package integration

import (
	"github.com/QUDUSKUNLE/shipping/internals/core/ports"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
)

type (
	ExternalRepository struct {}
	ExternalServicesHandler struct {
		external ports.ExternalPorts
	}
	ExternalServicesFacade struct {
		terminalService *domain.Terminal
	}
)

func OpenExternalConnection() *ExternalRepository {
	return &ExternalRepository{}
}

func ExternalServicesAdapter(externalPorts ports.ExternalPorts) *ExternalServicesHandler {
	return &ExternalServicesHandler{
		external: externalPorts,
	}
}

func NewExternalServicesFacade() *ExternalServicesFacade {
	return &ExternalServicesFacade{
		terminalService: &domain.Terminal{},
	}
}
