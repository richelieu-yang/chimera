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

// GetIpInfo
/*
PS: 仅支持IPV4，不支持国外IP解析.
*/
func (client *Client) GetIpInfo(ip string) (*IpInfo, error) {
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

	/* (1) 特殊处理: ip为内网地址 */
	// ip为内网地址情况下的响应json例子: {"status":"1","info":"OK","infocode":"10000","province":"局域网","city":[],"adcode":[],"rectangle":[]}
	field := jsonKit.GetStringField(jsonData, "province")
	if field == "局域网" {
		return &IpInfo{
			Province:  field,
			City:      "",
			Adcode:    "",
			Rectangle: "",
		}, nil
	}

	// 外网ip的响应json例子: {"status":"1","info":"OK","infocode":"10000","province":[],"city":[],"adcode":[],"rectangle":[]}

	/* (2) 正常处理 */
	resp := &IpResponse{}
	if err := jsonKit.Unmarshal(jsonData, resp); err != nil {
		return nil, errorKit.Wrap(err, "Fail to unmarshal")
	}
	if err := resp.IsSuccess(); err != nil {
		return nil, err
	}
	return &resp.IpInfo, nil
}
