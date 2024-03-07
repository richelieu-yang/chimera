package reflectKit

import (
	"fmt"
	"testing"
)

func TestGetUnexportedFieldAddr(t *testing.T) {
	type Bean struct {
		id   uint
		flag bool
	}

	b := &Bean{
		id:   123,
		flag: true,
	}

	addr := GetUnexportedFieldAddr(b, "id")
	// 转换回原始类型，这里假设我们确切知道原始类型是什么
	i := *(*uint)(addr)
	fmt.Println(i) // 123

	addr1 := GetUnexportedFieldAddr(b, "flag")
	// 转换回原始类型，这里假设我们确切知道原始类型是什么
	f := *(*bool)(addr1)
	fmt.Println(f) // true
}
