package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Employee struct {
	Person   `json:",inline"`
	Position []string
}

func main() {
	e := Employee{
		Person: Person{
			Name: "John Doe",
			Age:  30,
		},
		Position: []string{"Software Engineer"},
	}
	fmt.Println(jsonKit.MarshalIndentToString(e, "", "    "))
}
