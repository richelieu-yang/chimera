package reqKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"time"
)

var defaultClient = NewClient()

// GetDefaultClient
/*
重用Client https://req.cool/zh/docs/tutorial/best-practices/#%e9%87%8d%e7%94%a8-client
	不要每次发请求都创建 Client，造成不必要的开销，通常可以复用同一 Client 发所有请求.

!!!: 不修改返回值的话，可以调用此方法；否则调用 NewClient.
*/
func GetDefaultClient() *req.Client {
	return defaultClient
}

func NewClient() *req.Client {
	client := req.C()

	// timeout（默认的2min太长了）
	client.SetTimeout(time.Second * 15)

	// 自动探测字符集并解码到 utf-8（默认就是启用）
	client.EnableAutoDecode()

	// 不验证非法的证书（默认验证）
	client.EnableInsecureSkipVerify()

	// 自定义 Marshal 和 Unmarshal
	api := jsonKit.GetAPI()
	client.SetJsonMarshal(api.Marshal).SetJsonUnmarshal(api.Unmarshal)

	return client
}
