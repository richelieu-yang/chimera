package gaodeKit

type (
	IpResponse struct {
		BaseResponse

		IpInfo
	}

	IpInfo struct {
		// Province 省份名称
		Province string `json:"province"`

		// City 城市名称
		City string `json:"city"`

		// Adcode 城市的adcode编码
		Adcode string `json:"adcode"`

		// Rectangle 所在城市矩形区域范围
		Rectangle string `json:"rectangle"`
	}
)
