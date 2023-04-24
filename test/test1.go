package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/urlKit"
)

func main() {
	m := map[string]string{
		"a": "test",
		"b": "测试",
	}
	fmt.Println(urlKit.ToQueryString(m)) // a=test&b=%E6%B5%8B%E8%AF%95
}
