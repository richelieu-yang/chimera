// Package netKit
package netKit

import (
	"net"
	"time"
)

var (
	DialTimeout func(network, address string, timeout time.Duration) (net.Conn, error) = net.DialTimeout
)

// CanDial
/*
@return 返回值如果为nil，说明能dial成功
*/
func CanDial(addr string, timeout time.Duration) error {
	conn, err := DialTimeout("tcp", addr, timeout)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
