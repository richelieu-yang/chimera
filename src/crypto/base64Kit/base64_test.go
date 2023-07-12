package base64Kit

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	str := "\u0000\u0000>\u0000\u0000?\u0000~!@陈192.168.9.254圆圆#$%^&*()_+-=~!2`1234567890-=<>,./:\";''"

	{
		// base64.StdEncoding
		str1 := EncodeStringToString(str)
		// !!!: 生成的base64字符串中有'+'、'/'、'='
		fmt.Println(str1)
		str2, err := DecodeStringToString(str1)
		if err != nil {
			panic(err)
		}
		fmt.Println(str2)
		if str != str2 {
			panic("str != str2")
		}
		fmt.Println("str == str2")
	}

	fmt.Println("**************************************")

	{
		// base64.RawURLEncoding
		str1 := EncodeStringToString(str, WithEncoding(base64.RawURLEncoding))
		fmt.Println(str1)
		str2, err := DecodeStringToString(str1, WithEncoding(base64.RawURLEncoding))
		if err != nil {
			panic(err)
		}
		fmt.Println(str2)
		if str != str2 {
			panic("str != str2")
		}
		fmt.Println("str == str2")
	}

	fmt.Println("**************************************")
	{
		// base64.URLEncoding + base64.NoPadding
		r := base64.NoPadding
		str1 := EncodeStringToString(str, WithEncoding(base64.URLEncoding), WithPadding(&r))
		fmt.Println(str1)
		str2, err := DecodeStringToString(str1, WithEncoding(base64.RawURLEncoding))
		if err != nil {
			panic(err)
		}
		fmt.Println(str2)
		if str != str2 {
			panic("str != str2")
		}
		fmt.Println("str == str2")
	}
}
