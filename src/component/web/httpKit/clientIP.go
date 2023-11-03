package httpKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"net"
	"net/http"
	"strings"
)

// GetClientIP
/*
!!!: 请注意，恶意用户可以创建伪造的 X-REAL-IP 和 X-FORWARDED-FOR 标头。因此，在处理这些头部信息时需要谨慎。如果可能，应该使用一种方法来验证IP地址的真实性。

流程:
(1) 这个函数首先尝试获取X-Real-IP头部信息，如果这个头部信息存在并且是一个有效的IP地址，那么就返回这个IP。
(2) 如果X-Real-IP头部信息不存在或者不是一个有效的IP地址，那么函数会尝试获取X-Forwarded-For头部信息，这个头部信息包含了一个IP地址列表，每个经过的代理服务器都会添加一个IP。
(3) 如果X-Forwarded-For头部信息中存在有效的IP地址，那么就返回这个IP。
(4) 如果以上两个头部信息都不存在有效的IP地址，那么函数会尝试获取RemoteAddr，这个是Web服务器从其接收连接并将响应发送到的实际物理IP地址。
*/
func GetClientIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forwarded-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", errorKit.Wrap(err, "Fail to get client ip with RemoteAddr(%s)", r.RemoteAddr)
	}
	if net.ParseIP(host) != nil {
		return host, nil
	}
	return "", errorKit.New("Fail to get client ip with host(%s)", host)
}
