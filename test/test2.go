package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/web/httpClientKit"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, data, _ := httpClientKit.Get("http://127.0.0.1:9942/ws/api/suicide")
			fmt.Println(string(data))
		}(i)
	}
	wg.Wait()
}
