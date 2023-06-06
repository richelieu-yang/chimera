package mqKit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
)

// NewProducer
/*
PS: In most case, you don't need to create many producers, singletion pattern is more recommended.

@param clientLogPath 客户端日志（blank则输出到控制台）
*/
func NewProducer() (rmq_client.Producer, error) {
	if config == nil {
		return nil, NotSetupError
	}

	endpoint := sliceKit.Join(config.Endpoints, ";")

	producer, err := rmq_client.NewProducer(&rmq_client.Config{
		Endpoint:    endpoint,
		Credentials: config.Credentials,
	})
	if err != nil {
		return nil, err
	}
	// Start
	if err := producer.Start(); err != nil {
		return nil, err
	}
	return producer, nil
}
