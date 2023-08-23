package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/cpuKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/web/httpClientKit"
	"time"
)

func init() {
	cpuKit.SetUp()
}

func main() {
	code, data, err := httpClientKit.Get("http://127.0.0.1/ping", httpClientKit.WithTimeout(time.Second*2))
	if err != nil {
		err = errorKit.Wrap(err, "1")
		fmt.Println(httpClientKit.IsTimeoutError(err))
		panic(err)
	}
	fmt.Println(code, string(data))

	//req, err := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1/ping", nil)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//
	//client := http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	if errorKit.Is(err, context.DeadlineExceeded) {
	//		fmt.Println("ccc")
	//	}
	//	fmt.Println("Error:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//if resp.StatusCode == http.StatusOK {
	//	fmt.Println("Request successful")
	//} else {
	//	fmt.Println("Request failed with status:", resp.Status)
	//}
}
