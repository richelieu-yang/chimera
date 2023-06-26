// Package netKit
package netKit

import (
	"net"
	"time"
)

// Dial
/*
判断指定ip和port能否连通.

PS: 超时时间（timeout）默认为300ms.
*/
func Dial(addr string) (bool, error) {
	return DialTimeout(addr, time.Millisecond*500)
}

// DialTimeout
/*
判断指定ip和port能否连通（在指定时限内）.
*/
func DialTimeout(addr string, timeout time.Duration) (bool, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	return true, nil
}
