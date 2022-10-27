package httpKit

import (
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/urlKit"
	"io"
	"net/http"
	"strings"
)

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
		if req.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}
	}

	/* host */
	host := url.Host
	if strKit.IsEmpty(host) {
		host = req.Host
	}

	/* path */
	path := url.Path

	return strKit.Format("%s://%s%s", scheme, host, path)
}
