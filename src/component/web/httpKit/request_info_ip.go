package httpKit

import (
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"net"
	"net/http"
	"strings"
)

var (
	// GetRequestPublicIp 获取http请求的 "公网ip".
	/*
		涉及的请求头:
		(1) "X-Forwarded-For"
		(2) "X-Real-Ip"
	*/
	GetRequestPublicIp func(req *http.Request) string = netutil.GetRequestPublicIp
)

// GetRemoteIP 获取客户端IP地址（客户端的远程IP地址）.
/*
PS: 参考 gin's Context.RemoteIP().

e.g.
当客户端通过代理服务器连接时，RemoteIP() 返回代理服务器的 IP 地址
*/
func GetRemoteIP(req *http.Request) string {
	//ctx := &gin.Context{
	//	Request: req,
	//}
	//return ctx.RemoteIP()

	ip, _, err := net.SplitHostPort(strings.TrimSpace(req.RemoteAddr))
	if err != nil {
		return ""
	}
	return ip
}

// GetClientIP
/*
PS: 参考 gin's Context.ClientIP().
*/
func GetClientIP(req *http.Request) string {
	ip := GetClientIPFromHeader(req)
	if strKit.IsEmpty(ip) {
		ip = GetRemoteIP(req)
	}
	return ip
}

// GetClientIPFromHeader
/*
PS: 参考 gin's Context.ClientIP().
*/
func GetClientIPFromHeader(req *http.Request) string {
	//ctx := &gin.Context{
	//	Request: req,
	//}
	//return ctx.ClientIP()

	remoteIPHeaders := []string{"X-Forwarded-For", "X-Real-IP"}
	for _, headerName := range remoteIPHeaders {
		ip, valid := validateHeader(GetHeader(req.Header, headerName))
		if valid {
			return ip
		}
	}
	return ""
}

func validateHeader(header string) (clientIP string, valid bool) {
	if header == "" {
		return "", false
	}

	items := strings.Split(header, ",")
	for i := len(items) - 1; i >= 0; i-- {
		ipStr := strings.TrimSpace(items[i])
		ip := net.ParseIP(ipStr)
		if ip == nil {
			break
		}

		// X-Forwarded-For is appended by proxy
		// Check IPs in reverse order and stop when find untrusted proxy
		if i == 0 {
			return ipStr, true
		}
	}
	return "", false
}
