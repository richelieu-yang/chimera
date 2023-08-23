package main

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/cpuKit"
	"github.com/richelieu-yang/chimera/v2/src/web/httpClientKit"
	"net/http"
	"time"
)

func init() {
	cpuKit.SetUp()
}

func main() {
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*2)

	req, err := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1/ping", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("timeout: ", httpClientKit.IsTimeoutError(err))
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Request successful")
	} else {
		fmt.Println("Request failed with status:", resp.Status)
	}
}
