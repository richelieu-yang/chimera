package httpClientKit

import "time"

type (
	ClientConfig struct {
		// Timeout 请求的超时时间
		Timeout time.Duration

		InsecureSkipVerify bool
	}
)
