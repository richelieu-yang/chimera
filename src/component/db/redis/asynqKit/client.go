package asynqKit

import (
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/richelieu-yang/chimera/v2/src/component/db/redis/redisKit"
)

type optImpl struct {
	asynq.RedisConnOpt

	uc redis.UniversalClient
}

func (impl *optImpl) MakeRedisClient() interface{} {
	return impl.uc
}

// NewAsynqClient
/*
!!!: 必需先setup redis.
*/
func NewAsynqClient() (*asynq.Client, error) {
	client, err := redisKit.GetClient()
	if err != nil {
		return nil, err
	}
	impl := &optImpl{
		uc: client.GetUniversalClient(),
	}

	return asynq.NewClient(impl), nil
}
