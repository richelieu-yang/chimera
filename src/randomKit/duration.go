package randomKit

import (
	"github.com/gogf/gf/v2/util/grand"
	"time"
)

var (
	// Duration
	/*
		@return a random time.Duration between min and max: [min, max]
	*/
	Duration func(min, max time.Duration) time.Duration = grand.D
)
