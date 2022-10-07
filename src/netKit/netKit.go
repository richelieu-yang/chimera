// Package netKit
/**
 * 参考："net"
 */
package netKit

import (
	"net"
	"time"
)

// CanDial
/*
判断指定ip和port能否连通.

PS: 超时时间（timeout）默认为300ms.
*/
func CanDial(ip string, port int) (bool, error) {
	return CanDialWithTimeout(ip, port, time.Millisecond*300)
}

// CanDialWithTimeout
/*
判断指定ip和port能否连通（在指定时限内）.
*/
func CanDialWithTimeout(ip string, port int, timeout time.Duration) (bool, error) {
	address := JoinHostPort(ip, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	return true, nil
}
