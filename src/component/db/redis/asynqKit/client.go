package asynqKit

import "github.com/hibiken/asynq"

func NewClient() {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
}
