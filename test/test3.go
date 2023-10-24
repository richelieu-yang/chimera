package main

import (
	"fmt"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
)

func main() {
	//data := grand.B(1024)
	//fmt.Println(string(data))

	tmp := grand.Digits(int(dataSizeKit.MiB))
	fmt.Println(tmp)

	_ = fileKit.WriteToFile([]byte(tmp), "tmp.txt", 0644)

	//fmt.Println(jsonKit.MarshalToString(nil)) // null <nil>
}
