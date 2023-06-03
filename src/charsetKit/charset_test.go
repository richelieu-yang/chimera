package charsetKit

import (
	"fmt"
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
