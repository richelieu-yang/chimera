package componentKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/database/redisKit"
	"github.com/sirupsen/logrus"
)

var client *redisKit.Client

// InitializeRedisComponent 初始化Redis组件（可选）
func InitializeRedisComponent() error {
	config, err := GetRedisConfig()
	if err != nil {
		return err
	}
	if config == nil {
		return errorKit.Simple("config == nil")
	}

	client, err = redisKit.NewClient(config)
	if err == nil {
		logrus.Info("[COMPONENT, REDIS] Initialize successfully.")
	}
	return err
}

// GetRedisClient
/*
@return 如果Redis组件未被初始化，将返回nil
*/
func GetRedisClient() *redisKit.Client {
	return client
}
