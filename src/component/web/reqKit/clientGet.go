package reqKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/urlKit"
)

// SimpleGet
/*
!!!: 第2个返回值==nil的情况下，不需要手动关闭第1个返回值.
*/
func (c *Client) SimpleGet(url string, queryParams map[string][]string) (resp *req.Response, err error) {
	url, err = urlKit.PolyfillUrl(url, queryParams)
	if err != nil {
		return
	}

	for i := 0; i < c.maxRetryTimes; i++ {
		resp = c.Client.Get(url).Do()
		err = resp.Err
		if err != nil {
			break
		}
	}
	return resp, err
}

func (c *Client) Get(url string, queryParams map[string][]string) (code int, data []byte, err error) {
	var resp *req.Response
	resp, err = c.SimpleGet(url, queryParams)
	if err != nil {
		return
	}
	// 不需要手动关闭 resp
	//defer resp.Body.Close()

	code = resp.StatusCode
	data = resp.Bytes()
	if !resp.IsSuccessState() {
		err = errorKit.New("not success state(%d)", code)
	}
	return
}
