package httpClientKit

import "time"

var (
	// DefaultTimeout 发送请求的默认超时时间.
	/*
		PS:
		(1) 个人实测，对于 http.Client 结构体，Timeout 默认为30s.
		(2) e.g. yozo的网访问谷歌必定超时.
	*/
	DefaultTimeout = time.Second * 30
)
