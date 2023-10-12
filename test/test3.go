package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/mapKit"
)

func main() {
	m := map[string]interface{}{
		"a": 0,
	}

	err := mapKit.SetSafelyAndHandleOldValue(m, "a", 1, func(v interface{}) error {
		fmt.Println(v)
		return nil
	})
	fmt.Println(err)
	fmt.Println(m)
}
