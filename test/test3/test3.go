package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/config/confKit"
	"github.com/richelieu-yang/chimera/v2/src/serialize/json/jsonKit"
)

func main() {
	confKit.LoadFromJsonText()

	m := map[interface{}]interface{}{
		"0": 3.1415926,
	}

	data, err := jsonKit.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	m1 := map[interface{}]interface{}{}
	if err := jsonKit.Unmarshal(data, &m1); err != nil {
		panic(err) // panic: unsupported map key type: interface {}
	}
	fmt.Println(m1)
}
