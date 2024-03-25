package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

func main() {
	src := "/Users/richelieu/Downloads/魔都精兵的奴隶_副本/[Sakurato] Mato Seihei no Slave [11][AVC-8bit 1080p AAC][CHS].mp4"
	//dest := "/Users/richelieu/Downloads/02.mp4"
	fmt.Println(fileKit.RenameInSameDir(src, "11.mp4"))
}
