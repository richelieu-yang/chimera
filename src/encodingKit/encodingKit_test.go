package encodingKit

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	ori := "测试123~！@#￥%……&*（）——+-="

	// UTF-8 => GBK
	str, _, err := Utf8ToGbkString(ori)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GBK: [%s].\n", str)
	// GBK => UTF-8
	str, _, err = GbkToUtf8String(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("UTF-8: [%s].\n", str)
	if str != ori {
		t.Fatal("第1次测试失败！")
		return
	}

	// UTF-8 => GBK
	data, err := Utf8ToGbk([]byte(str))
	if err != nil {
		panic(err)
	}
	str = string(data)
	fmt.Printf("GBK: [%s].\n", str)

	// GBK => UTF-8
	data, err = GbkToUtf8([]byte(str))
	if err != nil {
		panic(err)
	}
	str = string(data)
	fmt.Printf("UTF-8: [%s].\n", str)
	if str != ori {
		t.Fatal("第2次测试失败！")
		return
	}
}
