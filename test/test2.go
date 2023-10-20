package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/reflectKit"
	"reflect"
)

func main() {
	v := reflectKit.ValueOf(111)
	fmt.Println(v.Kind() == reflect.Int) // true
}
