package main

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/richelieu-yang/chimera/v2/src/json/sonicKit"
)

func main() {
	m := map[string]interface{}{
		"a":    1,
		"b":    2,
		"1":    3,
		"111":  6,
		"1111": 6,
		"1112": 6,
		"d":    4,
	}
	for i := 0; i < 1000000; i++ {
		str, _ := sonicKit.MarshalToStringByAPI(sonic.ConfigDefault, m)
		if str != `{"1":3,"111":6,"1111":6,"1112":6,"a":1,"b":2,"d":4}` {
			panic(str)
		}
		fmt.Println(str)
	}
}
