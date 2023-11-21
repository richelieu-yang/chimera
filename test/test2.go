package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/component/web/reqKit"
)

func main() {
	status, data, err := reqKit.Post("https://127.0.0.1/test", nil, nil)
	fmt.Println(status)
	fmt.Println(string(data))
	fmt.Println(err)
}
