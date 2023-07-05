package main

import (
	"fmt"
	"github.com/spf13/cast"
	"strconv"
)

func main() {
	fmt.Println(strconv.Atoi("08"))  // 8 <nil>
	fmt.Println(strconv.Atoi("0xg")) // 0 strconv.Atoi: parsing "0xg": invalid syntax

	// 0开头是"8进制数"，8进制数不应该出现"8"，所以转换失败
	fmt.Println(cast.ToIntE("08")) // 0 unable to cast "08" of type string to int64
	// 0x开头是"16进制数"，16进制数不应该出现"g"，所以转换失败
	fmt.Println(cast.ToIntE("0xg")) // 0 unable to cast "0xg" of type string to int64
}
