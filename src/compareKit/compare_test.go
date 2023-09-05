package compareKit

import (
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	type bean struct {
		Name   string
		Lovers []string
	}
	b0 := &bean{
		Name:   "张三",
		Lovers: []string{"李四"},
	}
	b1 := &bean{
		Name:   "张三",
		Lovers: []string{"李四"},
	}
	fmt.Println(Equal(b0, b1))
}

type wrapper struct {
	Bean *bean
}

type bean struct {
	Name   string
	Lovers []string
}

func (b bean) Equal(b1 bean) bool {
	return b.Name == b1.Name
}

// TestEqual1 结构体实现了: (T) Equal(T) bool 或者 (T) Equal(I) bool
/*
!!!:
(1) receiver建议为 值类型.
*/
func TestEqual1(t *testing.T) {
	w0 := &wrapper{
		Bean: &bean{
			Name:   "张三",
			Lovers: []string{"李四"},
		},
	}
	w1 := &wrapper{
		Bean: &bean{
			Name:   "张三",
			Lovers: []string{"李4"},
		},
	}
	fmt.Println(Equal(w0, w1))
}
