package netKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"time"
)

const (
	DefaultHttpPort  = 80
	DefaultHttpsPort = 443

	// MaxPort 65535 == 0xFFFF
	MaxPort = 0xFFFF
)

var (
	NoUsablePortError = errorKit.New("no usable port")
)

// IsValidPort 是否为有效的端口？（根据值范围判断）
/*
参考:
Java，hutool中的NetUtil.isValidPort()
Linux端口分配: https://blog.csdn.net/zh2508/article/details/104888743

0 			不使用
1–1023 		系统保留,只能由root用户使用
1024—4999 	由客户端程序自由分配
5000—65535 	由服务器端程序自由分配（65535 = 2 ^ 16 - 1）
*/
func IsValidPort(port int) bool {
	return port > 0 && port <= MaxPort
}

// IsLocalPortAvailable 本地端口是否可用（即未被占用）？
// Deprecated: 某些绑定非127.0.0.1的端口无法被检测到.
/*
PS: 会优先判断端口是否有效（valid）.

参考:
Java，hutool中的NetUtil.isUsableLocalPort()
golang端口占用检测的使用	https://wenku.baidu.com/view/25716f5b01768e9951e79b89680203d8ce2f6af5.html
*/
func IsLocalPortAvailable(port int) bool {
	if !IsValidPort(port) {
		return false
	}

	// 能连通就说明端口被占用了
	addr := JoinHostnameAndPort("127.0.0.1", port)

	return CanDial(addr, time.Second*3) != nil
}
