package httpKit

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"github.com/richelieu42/chimera/v2/src/operationKit"
	"github.com/richelieu42/chimera/v2/src/urlKit"
	"io"
	"net/http"
	"strings"
)

// GetRequestRoute 获取请求的路由.
func GetRequestRoute(req *http.Request) string {
	return req.URL.Path
}

// GetProto
/*
@return "HTTP/1.0" || "HTTP/1.1" || ...
*/
func GetProto(req *http.Request) string {
	return req.Proto
}

// OverridePostRequestBody 覆盖POST请求的请求体.
func OverridePostRequestBody(req *http.Request, m map[string]string) {
	content := urlKit.ToQueryString(m)
	reader := strings.NewReader(content)

	// 下面2行代码二选一，都可以
	//req.Body = &Repeat{Reader: reader, Offset: 0}
	req.Body = io.NopCloser(reader)

	req.ContentLength = int64(len(content))
}

// GetRequestUrl 返回当前接口的url.
/*
PS: 不包括query数据.
*/
func GetRequestUrl(req *http.Request) string {
	url := req.URL

	/* scheme */
	scheme := url.Scheme
	if strKit.IsEmpty(scheme) {
		if websocket.IsWebSocketUpgrade(req) {
			scheme = operationKit.Ternary(req.TLS != nil, "wss", "ws")
		} else {
			scheme = operationKit.Ternary(req.TLS != nil, "https", "http")
		}
	}

	/* host */
	host := url.Host
	if strKit.IsEmpty(host) {
		host = req.Host
	}

	/* path */
	path := url.Path

	return fmt.Sprintf("%s://%s%s", scheme, host, path)
}
