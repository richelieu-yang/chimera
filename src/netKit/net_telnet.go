package netKit

import "github.com/duke-git/lancet/v2/netutil"

var (
	// IsTelnetConnected 检查能否telnet到主机.
	/*
		PS: 超时时间为5s.

		e.g.
			result1 := netutil.IsTelnetConnected("www.baidu.com", "80")
			result2 := netutil.IsTelnetConnected("www.baidu.com", "123")
			fmt.Println(result1)	// true
			fmt.Println(result2)	// false
	*/
	IsTelnetConnected func(hostname string, port string) bool = netutil.IsTelnetConnected
)
