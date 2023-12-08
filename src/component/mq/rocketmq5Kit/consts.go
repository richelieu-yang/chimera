package rocketmq5Kit

import "time"

const (
	// DefaultAwaitDuration (Consumer)maximum waiting time for receive func
	DefaultAwaitDuration = time.Second * 5

	// DefaultMaxMessageNum (Consumer)maximum number of messages received at one time
	DefaultMaxMessageNum int32 = 64

	// DefaultInvisibleDuration (Consumer)should > 20s
	DefaultInvisibleDuration = time.Second * 20
)
