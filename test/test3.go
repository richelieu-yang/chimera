package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func main() {
	data, err := fileKit.ReadFile("/Users/richelieu/GolandProjects/chimera/src/charsetKit/_gbk.txt")
	if err != nil {
		panic(err)
	}
	//data := []byte("test测试")

	simplifiedchinese.GBK

	r, err := chardet.NewTextDetector().DetectBest(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Charset, r.Language, r.Confidence)
}
