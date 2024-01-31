package netKit

import "github.com/duke-git/lancet/v2/netutil"

var (
	// IsPingConnected 检查能否ping通主机.
	/*
		e.g.
			result1 := netutil.IsPingConnected("www.baidu.com")
		    result2 := netutil.IsPingConnected("www.!@#&&&.com")
		    fmt.Println(result1)	// true
		    fmt.Println(result2)	// false
	*/
	IsPingConnected func(hostname string) bool = netutil.IsPingConnected
)
