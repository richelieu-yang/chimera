package base64Kit

import (
	"fmt"
	"os"
	"testing"
)

// 测试 []byte
func Test0(t *testing.T) {
	path := "/Users/richelieu/Desktop/a.JPG"

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	dest := Encode(data)
	if err != nil {
		panic(err)
	}

	src1, err := Decode(dest)
	if err != nil {
		panic(err)
	}

	if len(data) != len(src1) {
		panic("前后长度不一！")
	} else {
		fmt.Println("前后长度一样.")
	}
	if string(data) != string(src1) {
		panic("前后内容不一！")
	} else {
		fmt.Println("前后内容一样.")
	}
	fmt.Printf("-----------------------------")
}

// 测试 string
func Test1(t *testing.T) {
	a := "gqweuydfqwd强无敌群无！@#￥%……&*（）——+-=【】「」，。、《》？；'：、\"多"

	b := EncodeToString([]byte(a))
	a1, err := DecodeToString([]byte(b))
	if err != nil {
		panic(err)
	}
	if a != a1 {
		panic("前后内容不一！")
	} else {
		fmt.Println("前后内容一样.")
	}
	fmt.Printf("-----------------------------")
}
