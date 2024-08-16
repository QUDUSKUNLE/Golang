package integration

import (
	ports "github.com/QUDUSKUNLE/shipping/internals/core/ports"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
)

type ExternalServicesHandler struct {
	external ports.ExternalPorts
}

func ExternalServicesAdapter(externalPorts ports.ExternalPorts) *ExternalServicesHandler {
	return &ExternalServicesHandler{
		external: externalPorts,
	}
}

type ExternalServicesFacade struct {
	terminalService *domain.Terminal
}


func (externalServicesHandler *ExternalServicesHandler) NewExternalServicesFacade() *ExternalServicesFacade {
	return &ExternalServicesFacade{
		terminalService: &domain.Terminal{},
	}
}
