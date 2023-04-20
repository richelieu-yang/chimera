package httpClientKit

import (
	"crypto/tls"
	"github.com/richelieu42/chimera/v2/src/assertKit"
	"github.com/richelieu42/chimera/v2/src/urlKit"
	"io"
	"net/http"
)

func Get(url string, options ...Option) (int, []byte, error) {
	opts := loadOptions(options...)

	if err := assertKit.AssertHttpUrl(url); err != nil {
		return 0, nil, err
	}
	url = urlKit.AttachQueryParamsToUrl(url, opts.urlParams)

	client := &http.Client{
		Timeout: opts.timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !opts.safe,
			},
		},
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, data, nil
}
