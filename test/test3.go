package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{"John", 30}
	v := reflect.ValueOf(&p).Elem()
	f := v.FieldByName("name")
	if f.IsValid() && f.CanSet() {
		f.SetString("Doe")
	}
	fmt.Println(p)
}
