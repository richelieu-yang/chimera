package rocketmq5Kit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/jinzhu/copier"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
)

var (
	defaultCredentials = &credentials.SessionCredentials{
		AccessKey:    "",
		AccessSecret: "",
	}
)

func processBaseConfig(srcConfig *rmq_client.Config) (*rmq_client.Config, error) {
	if srcConfig == nil {
		return nil, errorKit.Simple("config == nil")
	}

	// 为了不修改传参 srcConfig 而新建了 config
	config := &rmq_client.Config{}
	if err := copier.Copy(config, srcConfig); err != nil {
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
