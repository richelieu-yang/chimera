package reqKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
)

func NewClient() *req.Client {
	client := req.C()

	// 自动探测字符集并解码到 utf-8（默认就是启用）
	client.EnableAutoDecode()

	// 不验证非法的证书（默认验证）
	client.EnableInsecureSkipVerify()

	// 自定义 Marshal 和 Unmarshal
	api := jsonKit.GetAPI()
	client.SetJsonMarshal(api.Marshal).SetJsonUnmarshal(api.Unmarshal)

	return client
}
