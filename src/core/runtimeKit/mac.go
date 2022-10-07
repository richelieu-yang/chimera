package runtimeKit

import (
	"gitee.com/richelieu042/go-scales/src/core/errorKit"
	"gitee.com/richelieu042/go-scales/src/core/strKit"
	"net"
)

// GetMacAddresses
/**
 * 参考：https://blog.csdn.net/chixielao6059/article/details/100860506
 *
 * MAC地址：主机网卡的物理地址.
 */
func GetMacAddresses() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	length := len(interfaces)
	if length == 0 {
		return nil, errorKit.Simple("length of interfaces is 0")
	}

	s := make([]string, 0, length)
	for _, inter := range interfaces {
		addr := inter.HardwareAddr.String()
		if strKit.IsNotEmpty(addr) {
			s = append(s, addr)
		}
	}
	if len(s) == 0 {
		return nil, errorKit.Simple("length of s is 0")
	}
	return s, nil
}
