package reqKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
)

func NewClient() *req.Client {
	client := req.C()

	// 自定义 Marshal 和 Unmarshal
	api := jsonKit.GetAPI()
	client.SetJsonMarshal(api.Marshal).SetJsonUnmarshal(api.Unmarshal)

	return client
}
