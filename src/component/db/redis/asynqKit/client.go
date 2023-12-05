package asynqKit

import (
	"github.com/hibiken/asynq"
	"github.com/richelieu-yang/chimera/v2/src/component/db/redis/redisKit"
)

type optImpl struct {
}

func (impl optImpl) MakeRedisClient() interface{} {
	client, _ := redisKit.GetClient()
	return client
}

func NewClient() (c *asynq.Client, err error) {
	redisKit.MustGetClient()

	opt := asynq.RedisClientOpt{Addr: "localhost:6379"}
	return asynq.NewClient(opt), nil
}
