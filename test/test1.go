package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/mapKit"
	"github.com/samber/lo"
)

func main() {

	lo.SliceToMap()

	s := mapKit.MapToSlice[string, string, string](map[string]string{"1": "a"}, func(key string, value string) string {
		return key + value
	})
	fmt.Println(s) // [1a]
}
