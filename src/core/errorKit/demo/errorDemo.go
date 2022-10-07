package main

import "os"

// PathError 自定义错误
type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + e.Path + e.Err.Error()
}

// demo1 返回实例
func demo1() error {
	return &PathError{"Op", "Path", nil}
}

// demo2 类型判断
func demo2() {
	_, err := os.Stat("a.txt")

	if err != nil {
		if e, ok := err.(*PathError); ok {
			println(e.Error())
		}
	}
}
