package ipKit

import (
	"github.com/duke-git/lancet/v2/netutil"
	"net"
)

var (
	// GetPublicIpInfo 获取公网ip信息.
	/*
		PS: 涉及发送http请求, 会有一定的耗时（得考虑内网环境）.

		e.g.
		info, err := ipKit.GetPublicIpInfo()
		if err != nil {
			panic(err)
		}
		fmt.Println(jsonKit.MarshalIndentToString(info, "", "    "))

		# Output
		{
		    "status": "success",
		    "country": "China",
		    "countryCode": "CN",
		    "region": "JS",
		    "regionName": "Jiangsu",
		    "city": "Suzhou",
		    "lat": 31.3093,
		    "lon": 120.602,
		    "isp": "Chinanet",
		    "org": "Chinanet JS",
		    "as": "AS4134 CHINANET-BACKBONE",
		    "query": "49.93.33.211"
		} <nil>
	*/
	GetPublicIpInfo func() (*netutil.PublicIpInfo, error) = netutil.GetPublicIpInfo

	// IsPublicIP 判断ip是否是公共ip.
	IsPublicIP func(IP net.IP) bool = netutil.IsPublicIP
)
