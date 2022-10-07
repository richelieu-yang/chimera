package timeKit

import (
	"net/http"
	"time"
)

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

// GetMachineTime 获取机器时间（本地时间；time.Local）.
func GetMachineTime() time.Time {
	return time.Now()
}

// GetNetworkTime
/*
获取网络时间.
PS: 获取不到，则返回机器时间.
*/
func GetNetworkTime() (time.Time, string, error) {
	var err error

	for _, source := range networkTimeSources {
		var t time.Time
		t, err = getNetworkTime(source)
		if err == nil {
			return t, source, nil
		}
	}
	return GetMachineTime(), "machine", err
}

func getNetworkTime(source string) (time.Time, error) {
	resp, err := http.Get(source)
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil {
		return time.Time{}, err
	}

	timeString := resp.Header.Get("Date")
	t, err := ParseStringToTime(string(NetworkFormat), timeString)
	if err != nil {
		return time.Time{}, err
	}
	return ConvertToLocalLocation(t), nil
}
