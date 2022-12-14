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

//// ParseStringToTime string => time.Time
///*
//PS: 采用本地时区(time zone).
//*/
//func ParseStringToTime(timeString, format string) (time.Time, error) {
//	return time.ParseStringToTimeWithLocation(format, timeString, time.Local)
//}
//
//func ParseStringToTimeWithLocation(timeString, format string, loc *time.Location) (time.Time, error) {
//	if loc == nil {
//		return time.Time{}, errorKit.Simple("loc is nil")
//	}
//	// 此处如果loc为nil，会panic("time: missing Location in call to Date").
//	return time.ParseStringToTimeWithLocation(format, timeString, loc)
//}

// GetSystemTime 获取系统时间（机器时间；本地时间；time.Local）.
func GetSystemTime() time.Time {
	return time.Now()
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
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()

	timeString := resp.Header.Get("Date")
	t, err := ParseStringToTime(string(NetworkFormat), timeString)
	if err != nil {
		return time.Time{}, err
	}
	return ConvertToLocalLocation(t), nil
}
