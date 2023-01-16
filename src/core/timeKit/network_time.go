package timeKit

import (
	"crypto/tls"
	"net/http"
	"time"
)

var networkTimeSources = []string{
	"http://www.ntsc.ac.cn",
	"http://www.taobao.com",
	"http://www.baidu.com",
	"http://www.360.cn",
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
		PS: 内网环境，可能直接返回error，也可能请求超时而返回error.
	*/
	client := &http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Get(source)
	if err != nil {
		return time.Time{}, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	timeString := resp.Header.Get("Date")
	t, err := ParseStringToTime(string(NetworkFormat), timeString)
	if err != nil {
		return time.Time{}, err
	}
	return ConvertToLocalLocation(t), nil
}
