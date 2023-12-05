package main

import "github.com/hibiken/asynq"

func main() {
	redisConnOpt := asynq.RedisClientOpt{
		Addr: "localhost:6379",
		// Omit if no password is required
		Password: "mypassword",
		// Use a dedicated db number for asynq.
		// By default, Redis offers 16 databases (0..15)
		DB: 0,
	}

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
}
