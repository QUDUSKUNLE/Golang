package ports

// Ports that connect external services or third party services
type ExternalPorts interface {
	TerminalRequest(endpoint ENDPOINTS, method METHOD, parameter any) (map[string]interface{}, error)
}
