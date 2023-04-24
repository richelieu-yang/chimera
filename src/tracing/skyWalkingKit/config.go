package skyWalkingKit

import "time"

type (
	Config struct {
		ServerAddr    string        `json:"serverAddr"`
		CheckInterval time.Duration `json:"checkInterval"`
	}
)
