package httpKit

import (
	"net"
	"net/http"
	"strings"
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
