package charsetKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
	"testing"
)

func TestConvert(t *testing.T) {
	//charset := "GBK"
	charset := "gbk"
	text := "`1234567890-=~!@#$%^&*()_+<>?,./;':\"[]\\{}|强无敌群无多我饿去当前文档"

	tmp, err := UTF8To(charset, text)
	if err != nil {
		panic(err)
	}
	fmt.Println("GBK:", tmp)

	tmp, err = ToUTF8(charset, tmp)
	if err != nil {
		panic(err)
	}
	fmt.Println("UTF-8:", tmp)

	if tmp != text {
		panic("not equal")
	}
}

// 获取txt文本的编码
func TestDetermineEncoding(t *testing.T) {
	gbkStr, err := fileKit.ReadFileToString("_gbk.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("gbkStr:", gbkStr)
	charset, err := Detect([]byte(gbkStr))
	if err != nil {
		panic(err)
	}
	fmt.Println(charset)

	fmt.Println("------")

	utf8Str, err := ToUTF8(charset, gbkStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("utf8Str:", utf8Str)
	fmt.Println(DetermineEncoding([]byte(utf8Str), ""))
}
