package consulKit

import "github.com/hashicorp/consul/api"

func GetAllServices() (map[string]*api.AgentService, error) {
	if innerClient == nil {
		return nil, NotSetupError
	}

	return innerClient.Agent().Services()
}

func GetAllServicesWithFilter(filter string) (map[string]*api.AgentService, error) {
	if innerClient == nil {
		return nil, NotSetupError
	}

	return innerClient.Agent().ServicesWithFilter(filter)
}

func GetAllServicesWithFilterOpts(filter string, q *api.QueryOptions) (map[string]*api.AgentService, error) {
	if innerClient == nil {
		return nil, NotSetupError
	}

	return innerClient.Agent().ServicesWithFilterOpts(filter, q)
}

func GetHealthServices(service, tag string, q *api.QueryOptions) ([]*api.ServiceEntry, *api.QueryMeta, error) {
	if innerClient == nil {
		return nil, nil, NotSetupError
	}

	return innerClient.Health().Service(service, tag, true, q)
}

func GetHealthServicesMultipleTags(service string, tags []string, q *api.QueryOptions) ([]*api.ServiceEntry, *api.QueryMeta, error) {
	if innerClient == nil {
		return nil, nil, NotSetupError
	}

	return innerClient.Health().ServiceMultipleTags(service, tags, true, q)
}
