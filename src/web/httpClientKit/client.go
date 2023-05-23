package httpClientKit

import (
	"crypto/tls"
	"net/http"
	"time"
)

func newHttpClient(timeout time.Duration, safe bool) *http.Client {
	return &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !safe,
			},
		},
	}
}
