package timeKit

import (
	"crypto/tls"
	"net/http"
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
		PS:
		(1) 此处不使用 httpClientKit，原因: 避免import cycle
		(2) 内网环境，可能的情况: (a)直接返回error; (b)请求超时而返回error.
	*/
	client := &http.Client{
		Timeout: time.Second * 3,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	req, err := http.NewRequest("GET", source, nil)
	if err != nil {
		return time.Time{}, err
	}
	resp, err := client.Do(req)
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
