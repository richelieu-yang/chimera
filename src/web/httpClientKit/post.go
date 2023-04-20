package httpClientKit

import (
	"github.com/richelieu42/chimera/v2/src/assertKit"
	"github.com/richelieu42/chimera/v2/src/urlKit"
	"io"
	"net/http"
	"strings"
)

func Post(url string, options ...Option) (int, []byte, error) {
	resp, err := PostForResponse(url, options...)
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

// PostForResponse
/*
@return !!!: 第一个返回值如果不为nil的话，一般来说需要手动调用 "resp.Body.Close()".
*/
func PostForResponse(url string, options ...Option) (*http.Response, error) {
	opts := loadOptions(options...)

	// url
	if err := assertKit.AssertHttpUrl(url); err != nil {
		return nil, err
	}
	url = urlKit.AttachQueryParamsToUrl(url, opts.urlParams)

	// body
	body := strings.NewReader(urlKit.ToBodyString(opts.postParams))

	// req
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", "charset=utf-8")

	// client
	client := newHttpClient(opts.timeout, opts.safe)

	return send(client, req)
}
