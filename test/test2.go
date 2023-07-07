package main

import (
	"fmt"
	"github.com/bytedance/sonic"
)

func main() {
	m := map[string]interface{}{
		"a": nil,
		"b": 0,
		"c": "",
	}
	fmt.Println(sonic.ConfigFastest.MarshalToString(m))
}
