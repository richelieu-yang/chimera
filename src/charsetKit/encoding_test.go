package charsetKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"testing"
)

func TestIsGBK(t *testing.T) {
	{
		data, err := fileKit.ReadFile("_gbk.txt")
		if err != nil {
			panic(err)
		}
		fmt.Println("IsGBK:", IsGBK(data))   // IsGBK: true
		fmt.Println("IsUTF8:", IsUTF8(data)) // IsUTF8: false
	}

	fmt.Println("------")

	{
		data, err := fileKit.ReadFile("_utf8.txt")
		if err != nil {
			panic(err)
		}
		fmt.Println("IsGBK:", IsGBK(data))   // IsGBK: false
		fmt.Println("IsUTF8:", IsUTF8(data)) // IsUTF8: true
	}
}
