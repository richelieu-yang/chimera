package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	logrusKit.MustSetUp(nil)

	// 创建一个HTTP客户端，设置超时时间为5秒
	client := http.Client{
		//Timeout: 5 * time.Second,
	}

	// 发起GET请求
	logrus.Info(0)
	resp, err := client.Get("http://127.0.0.1/ping")
	logrus.Info(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 处理响应
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Request successful")
	} else {
		fmt.Println("Request failed with status:", resp.Status)
	}
}
