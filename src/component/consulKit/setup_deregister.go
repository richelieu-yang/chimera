package consulKit

// Deregister 服务注销.
func Deregister(serviceID string) error {
	if innerClient == nil {
		return NotSetupError
	}

	return innerClient.Agent().ServiceDeregister(serviceID)
}
