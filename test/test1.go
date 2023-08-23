package main

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/cpuKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"net/http"
	"time"
)

func init() {
	cpuKit.SetUp()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1/ping", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if errorKit.Is(err, context.DeadlineExceeded) {
			fmt.Println("ccc")
		}
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Request successful")
	} else {
		fmt.Println("Request failed with status:", resp.Status)
	}
}
