package timeKit

import (
	"crypto/tls"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"net/http"
	"time"
)

// networkTimeSources 网络时间的来源s
var networkTimeSources = []string{
	"https://www.bilibili.com/",
	"https://www.baidu.com/",
	"https://cn.bing.com/",
	"http://www.ntsc.ac.cn/",
	"https://www.taobao.com/",
	"https://www.360.cn/",
	"https://www.google.com/",
	"https://www.kingsoft.com/",
	"https://www.yozosoft.com/",
}

// GetNetworkTime 获取网络时间.
/*
PS: 获取不到的话，返回机器时间.
*/
func GetNetworkTime() (time.Time, string, error) {
	type bean struct {
		source string
		t      time.Time
	}

	var timeout = time.Second * 10
	var ch = make(chan *bean, len(networkTimeSources))

	// 起多个goroutine同时获取网络时间，只要有一个成功获取到，此方法就返回值
	for _, source := range networkTimeSources {
		go func(url string) {
			t, err := getNetworkTimeBySource(url, timeout)
			if err != nil {
				return
			}
			ch <- &bean{
				source: url,
				t:      t,
			}
		}(source)
	}
	select {
	case b := <-ch:
		return b.t, b.source, nil
	case <-time.After(timeout):
		return time.Time{}, "", errorKit.New("timeout(%v)", timeout)
	}
}

func getNetworkTimeBySource(url string, timeout time.Duration) (time.Time, error) {
	/*
		超时时间设置的短一点，以防内网环境启动服务耗时太长.
		PS:
		(1) 此处不使用 httpClientKit，原因: 避免import cycle
		(2) 内网环境，可能的情况: (a)直接返回error; (b)请求超时而返回error.
	*/
	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	req, err := http.NewRequest("GET", url, nil)
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
