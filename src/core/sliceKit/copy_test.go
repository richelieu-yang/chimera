package sliceKit

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	m := map[string]interface{}{
		"Id":   0,
		"Name": "张三",
	}
	s := []interface{}{666, m}
	fmt.Println(s) // [666 map[Id:0 Name:张三]]

	s1 := Copy(s)
	s[0] = 999
	m["Name"] = "李四"
	fmt.Println(s)  // [999 map[Id:0 Name:李四]]
	fmt.Println(s1) // [666 map[Id:0 Name:李四]]
}

func TestDeepCopy(t *testing.T) {
	m := map[string]interface{}{
		"Id":   0,
		"Name": "张三",
	}
	s := []interface{}{666, m}
	s1, err := DeepCopy(s)
	if err != nil {
		panic(err)
	}

	s[0] = 999
	m["Name"] = "李四"

	fmt.Println(s)  // [999 map[Id:0 Name:李四]]
	fmt.Println(s1) // [666 map[Id:0 Name:张三]]
}
