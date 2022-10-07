package httpClientKit

import (
	"crypto/tls"
	"net/http"
)

// newClient
/*
参考: golang请求https跳过证书验证及获取证书信息 https://www.zhihuclub.com/86772.shtml

PS: 对于https请求，会跳过tls握手阶段里的证书校验.
*/
func newClient() *http.Client {
	return &http.Client{
		Timeout: DefaultTimeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
