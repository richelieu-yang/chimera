package reqKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/richelieu-yang/chimera/v2/src/urlKit"
	"time"
)

type Client struct {
	*req.Client

	// maxRetryTimes 最大重试次数
	maxRetryTimes int
}

func (c *Client) Get(url string, queryParams map[string][]string) (code int, data []byte, err error) {
	url, err = urlKit.PolyfillUrl(url, queryParams)
	if err != nil {
		return
	}

	var resp *req.Response
	for i := 0; i < c.maxRetryTimes; i++ {
		resp = c.Client.Get(url).Do()
		err = resp.Err
		if err != nil {
			break
		}
	}
	if err != nil {
		return 0, nil, err
	}

	// 不需要手动关闭
	//defer resp.Body.Close()

	code = resp.StatusCode
	data = resp.Bytes()
	if !resp.IsSuccessState() {
		err = errorKit.New("not success state(%d)", code)
	}
	return
}

var defaultClient = NewClient(3)

// GetDefaultClient
/*
重用Client https://req.cool/zh/docs/tutorial/best-practices/#%e9%87%8d%e7%94%a8-client
	不要每次发请求都创建 Client，造成不必要的开销，通常可以复用同一 Client 发所有请求.

!!!: 不修改返回值的话，可以调用此方法；否则调用 NewClient.
*/
func GetDefaultClient() *Client {
	return defaultClient
}

func NewClient(maxRetryTimes int) *Client {
	if maxRetryTimes <= 0 {
		maxRetryTimes = 3
	}

	client := req.C()

	// timeout（默认的2min太长了）
	client.SetTimeout(time.Second * 30)
	client.SetTLSHandshakeTimeout(time.Second * 30)

	// 自动探测字符集并解码到 utf-8（默认就是启用）
	client.EnableAutoDecode()

	// 不验证非法的证书（默认验证）
	client.EnableInsecureSkipVerify()

	// 自定义 Marshal 和 Unmarshal
	api := jsonKit.GetDefaultApi()
	client.SetJsonMarshal(api.Marshal).SetJsonUnmarshal(api.Unmarshal)

	return &Client{
		Client:        client,
		maxRetryTimes: maxRetryTimes,
	}
}
