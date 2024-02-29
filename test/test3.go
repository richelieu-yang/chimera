package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/langKit"
)

func main() {
	fmt.Println(langKit.S2T("0=无错误。\n1=程序退出中。\n\n1000=新的配置与旧的配置一致。\n\n2000=参数错误：%s。\n"))
}
