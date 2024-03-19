package consulKit

import (
	"github.com/hashicorp/consul/api"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

var (
	NotSetupError = errorKit.New("Haven’t been set up correctly")
)

var innerClient *api.Client

func MustSetUp(config *api.Config) {
	if err := SetUp(config); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(config *api.Config) (err error) {
	innerClient, err = NewClient(config)
	return
}

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

// Deregister 服务注销.
func Deregister(serviceID string) error {
	if innerClient == nil {
		return NotSetupError
	}

	return innerClient.Agent().ServiceDeregister(serviceID)
}

// Discover 服务发现.
func Discover() (map[string]*api.AgentService, error) {
	if innerClient == nil {
		return nil, NotSetupError
	}

	return innerClient.Agent().Services()
}

// DiscoverWithFilter 服务发现.
func DiscoverWithFilter(filter string) (map[string]*api.AgentService, error) {
	if innerClient == nil {
		return nil, NotSetupError
	}

	return innerClient.Agent().ServicesWithFilter(filter)
}

// DiscoverWithFilterOpts 服务发现.
func DiscoverWithFilterOpts(filter string, q *api.QueryOptions) (map[string]*api.AgentService, error) {
	if innerClient == nil {
		return nil, NotSetupError
	}

	return innerClient.Agent().ServicesWithFilterOpts(filter, q)
}
