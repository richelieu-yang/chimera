package httpClientKit

import "time"

type (
	ClientConfig struct {
		// Timeout http客户端的请求的超时时间
		Timeout time.Duration

		// InsecureSkipVerify http客户端，是否验证服务器端的 certificate chain and host name
		InsecureSkipVerify bool
	}
)
