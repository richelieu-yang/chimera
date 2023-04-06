package aesKit

import (
	"testing"
)

func Test(t *testing.T) {
	type arg struct {
		index     int
		plainText []byte
		key       []byte
	}

	args := []arg{
		{
			index:     0,
			plainText: []byte("test 测试"),
			key:       []byte("0123456789abcdef"),
		},
		{
			index:     1,
			plainText: []byte("q群文档群无多当前无多群无多群无多当前无多群无多群无多得去我多群无多当前文档群无多群得去我多群无多"),
			key:       []byte("0123456789abcdef"),
		},
	}

	for _, arg := range args {
		t.Logf("第[%d]次测试开始.......................................................", arg.index)
		base64Text, err := EncryptToString(arg.plainText, arg.key)
		if err != nil {
			t.Error(err)
			t.Fail()
			return
		}
		t.Logf("加密后: [%s].", base64Text)
		plainText1, err := DecryptToString([]byte(base64Text), arg.key)
		if err != nil {
			t.Error(err)
			t.Fail()
			return
		}
		t.Logf("解密后: [%s].", plainText1)
		if string(arg.plainText) != plainText1 {
			t.Error("加解密后的明文发生了变更！！！")
			t.Fail()
			return
		}
		t.Logf("第[%d]次测试通过.......................................................", arg.index)
	}
}
