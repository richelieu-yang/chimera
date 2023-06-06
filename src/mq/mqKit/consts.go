package mqKit

import "time"

const (
	// AwaitDuration (Consumer)maximum waiting time for receive func
	AwaitDuration = time.Second * 5

	// MaxMessageNum (Consumer)maximum number of messages received at one time
	MaxMessageNum int32 = 64

	// InvisibleDuration (Consumer)should > 20s
	InvisibleDuration = time.Second * 20
)
