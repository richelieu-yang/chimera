package gaodeKit

import "github.com/richelieu-yang/chimera/v2/src/ip/ipKit"

const (
	ipUrl = "https://restapi.amap.com/v3/ip"
)

// GetIp
/*
PS: 仅支持IPV4，不支持国外IP解析.
*/
func (client *Client) GetIp(ip string) (*IpInfo, error) {
	if err := ipKit.AssertIPv4(ip); err != nil {
		return nil, err
	}

}
