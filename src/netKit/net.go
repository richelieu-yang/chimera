package netKit

import "net"

var (
	// Listen 监听指定的网络地址.
	/*
		可能在以下几种情况下返回error：
		(1) 传参 network 不是有效值;
		(2) 传参 address 格式不正确;
		(3) 你试图监听的端口已经被其他进程占用; e.g. listen tcp 127.0.0.1:12345: bind: address already in use
		(4) 系统资源（如文件描述符）不足;
		(5) 如果你试图监听的端口号小于1024，而运行程序的用户没有足够的权限（通常需要 root 权限），函数将返回错误;
		(6) 如果网络出现问题（例如，网络接口不可用），函数将返回错误.

		@param network 网络类型，如: tcp、tcp4、tcp6、unix、unixpacket...
		@param address 监听的地址，格式为 ip:port，如果不指定 port，将由系统自动分配一个端口
	*/
	Listen func(network, address string) (net.Listener, error) = net.Listen

	// ResolveTCPAddr 将一个地址解析成TCP地址形式.
	/*
		@param network 网络类型， (1) 必须是: "tcp" || "tcp4" || "tcp6"
								(2) 如果是""，则默认为"tcp"
		@param address 地址
	*/
	ResolveTCPAddr func(network, address string) (*net.TCPAddr, error) = net.ResolveTCPAddr
)
