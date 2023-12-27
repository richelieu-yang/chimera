package netKit

import "net"

var (
	// Listen 监听指定的网络地址.
	/*
		PS:
		(1) 网络地址未被占用，返回的error == nil;
		(2)	网络地址已被占用，返回的error != nil.

		@param network 网络类型，如: tcp、tcp4、tcp6、unix、unixgram、unixpacket...
		@param address 监听的地址，格式为 ip:port，如果不指定 port，将由系统自动分配一个端口
	*/
	Listen func(network, address string) (net.Listener, error) = net.Listen
)
