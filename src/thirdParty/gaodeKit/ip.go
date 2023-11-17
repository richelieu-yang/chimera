package gaodeKit

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/reqKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/ip/ipKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
)

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

	_, jsonData, err := reqKit.Get(ipUrl, map[string][]string{
		"key": {client.key},
		"ip":  {ip},
	})
	if err != nil {
		return nil, err
	}

	resp := &IpResponse{}
	if err := jsonKit.Unmarshal(jsonData, resp); err != nil {
		return nil, errorKit.Wrap(err, "Fail to unmarshal")
	}
	if err := resp.IsSuccess(); err != nil {
		return nil, err
	}
	return &resp.IpInfo, nil
}
