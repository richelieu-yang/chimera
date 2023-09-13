package wsKit

import (
	"github.com/richelieu-yang/chimera/v2/src/web/httpKit"
	"net/http"
)

// PolyfillWebSocketRequest
/*
此函数是为了避免情况: 代理（e.g.Nginx）没有设置WebSocket穿透，导致WebSocket服务收到的WebSocket请求的header有问题.
*/
func PolyfillWebSocketRequest(r *http.Request) {
	httpKit.AddHeaderIfMissingIgnoreCase(r.Header, "Connection", "Upgrade")
	httpKit.AddHeaderIfMissingIgnoreCase(r.Header, "Upgrade", "websocket")
}
