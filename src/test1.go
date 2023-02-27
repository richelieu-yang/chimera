package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Person struct {
	name string
	age  int
}

func main() {
	p := &Person{
		name: "Alice",
		age:  25,
	}
	v := reflect.ValueOf(p).Elem()
	field := v.FieldByName("name")
	ptr := (*string)(unsafe.Pointer(field.UnsafeAddr()))

	// get
	fmt.Println(*ptr) // Alice

	// set
	*ptr = "Ccc"
	fmt.Println(p.name) // Ccc
}
