package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/richelieu42/go-scales/src/jsonKit"
)

func main() {
	m := map[string]interface{}{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
	}

	for i := 0; i < 10; i++ {
		fmt.Println(jsonKit.MarshalToStringWithJsoniterApi(jsoniter.ConfigCompatibleWithStandardLibrary, m))
	}
}
