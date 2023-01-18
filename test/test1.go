package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/mapKit"
)

func main() {
	fmt.Println(mapKit.Contains(map[string]interface{}(nil), "1"))

	v, err := mapKit.GetString(map[string]interface{}(nil), "1")
	fmt.Println(v, err)
}
