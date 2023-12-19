package compareKit

import (
	"fmt"
	"testing"
)

// 比较两个切片实例的内容是否相等.
func TestEqual(t *testing.T) {
	{
		s := []int{-1, 0, 1}
		s1 := []int{-1, 0, 1}
		fmt.Println(Equal(s, s1)) // true
	}

	{
		s := []int{-1, 0, 1}
		s1 := []int{0, 1, -1}
		fmt.Println(Equal(s, s1)) // false
	}
}

func TestEqual0(t *testing.T) {
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

// TestEqual1 要比较的结构体实现了: (T) Equal(T) bool 或者 (T) Equal(I) bool
/*
!!!:
(1) Equal方法的receiver为 指针类型 && 比较的是 结构体实例指针 的情况下，将返回true;
(2) Equal方法的receiver为 指针类型 && 比较的是 结构体实例 的情况下，将返回false（debug时，没有走到Equal方法）.
(3) Equal方法的receiver为 值类型类型，比较的无论是 结构体实例指针 还是 结构体实例，将返回true.
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
	fmt.Println(Equal(w0, w1)) // true
}

func TestEqual2(t *testing.T) {
	// 值都为nil，但类型不同
	var obj1 []string = nil
	var obj2 map[string][]string = nil
	fmt.Println(Equal(obj1, obj2)) // false

	// 值都为nil，且类型相同
	var obj3 map[string][]string = nil
	fmt.Println(Equal(obj2, obj3)) // true
}
