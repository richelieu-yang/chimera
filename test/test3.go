package main

import (
	"fmt"
	"github.com/spf13/cast"
	"strconv"
)

func main() {
	fmt.Println(cast.ToInt8E("07")) // 7 <nil>
	fmt.Println(cast.ToInt8E("08")) // 0 unable to cast "08" of type string to int64

	fmt.Println(strconv.Atoi("07")) // 7 <nil>
	fmt.Println(strconv.Atoi("08")) // 8 <nil>
}
