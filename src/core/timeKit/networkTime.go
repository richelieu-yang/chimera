package timeKit

import (
	"github.com/richelieu42/chimera/v2/src/web/httpClientKit"
	"time"
)

var networkTimeSources = []string{
	"https://www.baidu.com",
	"http://www.ntsc.ac.cn",
	//"http://www.taobao.com",
	//"http://www.360.cn",
}

// GetNetworkTime 获取网络时间.
/*
PS: 获取不到的话，返回机器时间.
*/
func GetNetworkTime() (time.Time, string, error) {
	var err error

	for _, source := range networkTimeSources {
		var t time.Time
		t, err = getNetworkTimeBySource(source)
		if err == nil {
			return t, source, nil
		}
	}
	return GetSystemTime(), "system", err
}

func getNetworkTimeBySource(source string) (time.Time, error) {
	/*
		超时时间设置的短一点，以防内网环境启动服务耗时太长.
		PS: 内网环境，可能的情况: (1) 直接返回error；(2) 请求超时而返回error.
	*/
	resp, err := httpClientKit.GetForResponse(source, httpClientKit.WithSafe(false), httpClientKit.WithTimeout(time.Second*6))
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	timeString := resp.Header.Get("Date")
	t, err := ParseTimeString(string(FormatNetwork), timeString)
	if err != nil {
		return time.Time{}, err
	}
	return ConvertToLocalLocation(t), nil
}
