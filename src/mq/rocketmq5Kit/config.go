package rocketmq5Kit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/richelieu42/chimera/src/copyKit"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/core/strKit"
)

type (
	Config struct {
		rmq_client.Config
		TopicToVerify string
	}
)

var (
	defaultCredentials = &credentials.SessionCredentials{
		AccessKey:    "",
		AccessSecret: "",
	}
)

// processConfig Consumer和Producer通用
func processConfig(baseConfig *rmq_client.Config) (*rmq_client.Config, error) {
	if baseConfig == nil {
		return nil, errorKit.Simple("config == nil")
	}

	// 深拷贝，为了不修改传参baseConfig
	config, err := copyKit.DeepCopy(baseConfig)
	if err != nil {
		return nil, err
	}

	if strKit.IsEmpty(config.Endpoint) {
		return nil, errorKit.Simple("config.Endpoint is empty")
	}

	config.ConsumerGroup = ""

	if config.Credentials == nil {
		config.Credentials = defaultCredentials
	}

	return config, nil
}
