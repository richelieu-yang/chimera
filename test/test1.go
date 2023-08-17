package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1/ping", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
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
