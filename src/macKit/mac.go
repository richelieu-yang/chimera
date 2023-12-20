package macKit

import "github.com/duke-git/lancet/v2/netutil"

var (
	// GetMacAddrs 获取mac地址列.
	/*
		e.g. 返回值
		[9a:e0:12:45:09:8f 9a:e0:12:45:09:8e 9a:e0:12:45:09:8d 9a:e0:12:45:09:6d 9a:e0:12:45:09:6e 9a:e0:12:45:09:6f 36:1a:80:3a:f6:80 36:1a:80:3a:f6:84 36:1a:80:3a:f6:88 f6:d4:88:80:13:31 f4:d4:88:80:13:31 36:1a:80:3a:f6:80 3e:89:30:b2:78:9d 3e:89:30:b2:78:9d 00:e0:4c:68:01:9b]
	*/
	GetMacAddrs func() []string = netutil.GetMacAddrs
)
