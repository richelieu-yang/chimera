package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/serialize/gobKit"
)

func main() {
	var m map[interface{}]interface{} = nil
	fmt.Println(gobKit.Marshal(m)) // [13 127 4 1 2 255 128 0 1 16 1 16 0 0 4 255 128 0 0] <nil>

	var obj interface{} = nil
	fmt.Println(gobKit.Marshal(obj)) // [] gob: cannot encode nil value
}
