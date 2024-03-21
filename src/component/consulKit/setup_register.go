package consulKit

import "github.com/hashicorp/consul/api"

// Register 服务注册.
func Register(service *api.AgentServiceRegistration) error {
	if innerClient == nil {
		return NotSetupError
	}

	return innerClient.Agent().ServiceRegister(service)
}

// RegisterOpts 服务注册.
func RegisterOpts(service *api.AgentServiceRegistration, opts api.ServiceRegisterOpts) error {
	if innerClient == nil {
		return NotSetupError
	}

	return innerClient.Agent().ServiceRegisterOpts(service, opts)
}
