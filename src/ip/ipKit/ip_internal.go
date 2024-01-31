package ipKit

import (
	"github.com/duke-git/lancet/v2/netutil"
	"net"
)

var (
	// GetInternalIp 获取内部ip.
	/*
		e.g.
		fmt.Println(ipKit.GetInternalIp()) // 172.20.10.4
	*/
	GetInternalIp func() string = netutil.GetInternalIp

	// IsInternalIP 判断ip是否是局域网ip.
	IsInternalIP func(IP net.IP) bool = netutil.IsInternalIP
)
