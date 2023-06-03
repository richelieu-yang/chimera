package rocketmq5Kit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"github.com/richelieu-yang/chimera/v2/src/copyKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

type (
	Config struct {
		rmq_client.Config
		TopicToVerify string
	}
)

var (
	defaultCredentials = &credentials.SessionCredentials{
		AccessKey:     "",
		AccessSecret:  "",
		SecurityToken: "",
	}
)

// processConfig Consumer和Producer通用
func processConfig(baseConfig *rmq_client.Config) (*rmq_client.Config, error) {
	if baseConfig == nil {
		return nil, errorKit.New("baseConfig == nil")
	}

	// 深拷贝（为了不修改传参baseConfig）
	config, err := copyKit.DeepCopy(baseConfig)
	if err != nil {
		return nil, err
	}

	if strKit.IsEmpty(config.Endpoint) {
		return nil, errorKit.New("config.Endpoint is empty")
	}
	config.ConsumerGroup = ""
	if config.Credentials == nil {
		config.Credentials = defaultCredentials
	}
	return config, nil
}
