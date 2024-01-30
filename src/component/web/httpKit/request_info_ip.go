package httpKit

import (
	"github.com/duke-git/lancet/v2/netutil"
	"net"
	"net/http"
	"strings"
)

var (
	// GetRequestPublicIp 获取http请求ip.
	GetRequestPublicIp func(req *http.Request) string = netutil.GetRequestPublicIp
)

// GetRemoteIP 获取客户端IP地址（客户端的远程IP地址）.
/*
PS: Copy from gin's Context.RemoteIP().

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

// GetClientIP 获取客户端IP地址（客户端的真实IP地址，但结果并不总是可靠的）.
/*
涉及的请求头:
(1) "X-Forwarded-For"
(2) "X-Real-Ip"
*/
var GetClientIP func(req *http.Request) string = netutil.GetRequestPublicIp
