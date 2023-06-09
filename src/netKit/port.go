package netKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"time"
)

const (
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
	return port >= 0 && port <= MaxPort
}

// IsLocalPortUsable 本地端口是否可用（即未被占用）？
// Deprecated: 某些绑定非127.0.0.1的端口无法被检测到.
/*
PS: 会优先判断端口是否有效（valid）.

参考:
Java，hutool中的NetUtil.isUsableLocalPort()
golang端口占用检测的使用	https://wenku.baidu.com/view/25716f5b01768e9951e79b89680203d8ce2f6af5.html
*/
func IsLocalPortUsable(port int) (bool, error) {
	if !IsValidPort(port) {
		return false, errorKit.New("port(%d) is invalid", port)
	}
	// 能连通就说明端口被占用了
	ok, _ := CanDialWithTimeout("127.0.0.1", port, time.Second*3)
	if ok {
		return false, errorKit.New("port(%d) is already in use", port)
	}
	return true, nil
}

// GetUsablePort 从 传参startPort开始(包括)向后找，返回一个可用的端口，或者直至端口超过上限65535.
/*
@param exceptivePorts 例外的端口（即返回的端口号不在exceptivePorts中）
*/
func GetUsablePort(startPort int, exceptivePorts ...int) (int, error) {
	if !IsValidPort(startPort) {
		return 0, errorKit.New("startPort(%d) is invalid", startPort)
	}

	for p := startPort; p <= MaxPort; p++ {
		if ok, _ := IsLocalPortUsable(p); ok {
			if !sliceKit.Contains(exceptivePorts, p) {
				return p, nil
			}
		}
	}
	return 0, NoUsablePortError
}
