package ipKit

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"net"
	"sync"
)

// GetLocalIp 获取本机ip
/*
Deprecated: 直接使用 GetOutboundIP()

参考:
系统性能数据gopsutil库 	https://topgoer.com/%E5%85%B6%E4%BB%96/%E7%B3%BB%E7%BB%9F%E6%80%A7%E8%83%BD%E6%95%B0%E6%8D%AEgopsutil%E5%BA%93.html
*/
func GetLocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return "", errorKit.Simple("fail to get local ip!")
}

var outboundIp string
var outboundErr error
var outboundOnce sync.Once

// GetOutboundIP 通过UDP获取本机ip(Get preferred outbound ip of this machine)
/*
参考:
系统性能数据gopsutil库 	https://topgoer.com/%E5%85%B6%E4%BB%96/%E7%B3%BB%E7%BB%9F%E6%80%A7%E8%83%BD%E6%95%B0%E6%8D%AEgopsutil%E5%BA%93.html
通过 UDP 获取本机 IP 		http://t.zoukankan.com/fousor-p-14874576.html

PS: 由于会申请一个UDP的端口，所以如果经常调用也会比较耗时的，这里如果需要可以将查询到的IP给缓存起来，性能可以获得很大提升.
*/
func GetOutboundIP() (string, error) {
	outboundOnce.Do(func() {
		callback := func() (string, error) {
			conn, err := net.Dial("udp", "8.8.8.8:80")
			if err != nil {
				return "", err
			}
			defer conn.Close()

			localAddr := conn.LocalAddr().(*net.UDPAddr)
			return localAddr.IP.String(), nil
		}
		outboundIp, outboundErr = callback()
	})
	return outboundIp, outboundErr
}
