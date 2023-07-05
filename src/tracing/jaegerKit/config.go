package jaegerKit

type (
	Config struct {
		Access bool `json:"access"`
		// Url Jaeger服务的url，e.g."http://localhost:14268/api/traces"
		Url string `json:"url"`
	}

	ResourceConfig struct {
		// ServiceName 服务名
		ServiceName string `json:"serviceName"`
		// Environment 环境
		Environment string `json:"environment"`
		// ID id
		ID int64 `json:"id"`
	}
)
