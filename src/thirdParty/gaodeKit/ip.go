package gaodeKit

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/request/reqKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/ip/ipKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/tidwall/gjson"
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

	/*
		局域网ip响应:	{"status":"1","info":"OK","infocode":"10000","province":"局域网","city":[],"adcode":[],"rectangle":[]}
		外网ip响应: 	{"status":"1","info":"OK","infocode":"10000","province":[],"city":[],"adcode":[],"rectangle":[]}
	*/
	field := jsonKit.GetField(jsonData, "province")
	if field.Type == gjson.String {
		if field.String() == "局域网" {
			return &IpInfo{
				Province:  "局域网",
				City:      "",
				Adcode:    "",
				Rectangle: "",
			}, nil
		} else {
			// continue
		}
	} else {
		if field.IsArray() && field.Raw == "[]" {
			return &IpInfo{
				Province:  "外网",
				City:      "",
				Adcode:    "",
				Rectangle: "",
			}, nil
		}
		return nil, errorKit.New("invalid response(%s)", string(jsonData))
	}

	/* 国内ip */
	resp := &IpResponse{}
	if err := jsonKit.Unmarshal(jsonData, resp); err != nil {
		return nil, errorKit.Wrap(err, "Fail to unmarshal")
	}
	if err := resp.IsSuccess(); err != nil {
		return nil, err
	}
	return &resp.IpInfo, nil
}
