package main

import "fmt"

func main() {
	var m = map[string]interface{}{}

	v, ok := m[""]
	fmt.Println(v, ok) // <nil> false

	m[""] = nil

	v, ok = m[""]
	fmt.Println(v, ok) // <nil> true
}
