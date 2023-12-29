package compareKit

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
)

// Equal 是否相等？
/*
!!!:
(0) 可能导致"卡死"（可能是死循环）的情况: 结构体内部的结构体实现了"(T) Equal(T) bool" 或者 "(T) Equal(I) bool";
(1) 如果传参结构体（或其内部的结构体）实现了 "(T) Equal(T) bool" 或者 "(T) Equal(I) bool"，
	且方法的receiver必须为 "值类型"，这样的话，无论比较 结构体实例 还是 结构体实例指针 都可以.

PS:
(1) 传参x、y都是 非nil的结构体实例指针，比较的是它们的内容是否相等;
(2) 传参x、y都是 切片实例，两者内容相同且顺序相同，将返回true；否则返回false;
(3) 如果传参分别为 结构体实例 和 结构体实例指针，将返回false;
(4) 传参x、y值都为nil的情况下，	(a) 类型不同，将返回 false;
							(b) 类型相同，将返回 true.

e.g.
	type person struct {
		Name string
		Age  int
	}

	p1 := &person{"Alice", 18}
	p2 := &person{"Bob", 20}
	p3 := &person{"Alice", 18}

	fmt.Println(compareKit.Equal(p1, p2)) // false
	fmt.Println(compareKit.Equal(p1, p3)) // true
*/
var Equal func(x, y interface{}, opts ...cmp.Option) bool = cmp.Equal

// Equal2 通过 reflect标准库 实现，比较2个结构体实例指针的内容是否相等.
/*
e.g.
	type person struct {
		Name string
		Age  int
	}

	p1 := &person{"Alice", 18}
	p2 := &person{"Bob", 20}
	p3 := &person{"Alice", 18}

	fmt.Println(reflect.DeepEqual(p1, p2)) // false
	fmt.Println(reflect.DeepEqual(p1, p3)) // true
*/
var Equal2 func(x, y any) bool = reflect.DeepEqual

// Diff 获取差异.
/*
!!!:
(0) 可能导致"卡死"（可能是死循环）的情况: 结构体内部的结构体实现了"(T) Equal(T) bool" 或者 "(T) Equal(I) bool";

@return 如果为""则说明两个传参一致.
*/
var Diff func(x, y interface{}, opts ...cmp.Option) string = cmp.Diff
