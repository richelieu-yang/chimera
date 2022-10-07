package main

import "fmt"

type Void struct{}

func main() {
	var member Void

	// New empty set
	set := make(map[string]Void)
	// Add
	set["Foo"] = member
	// Loop
	for ele := range set {
		fmt.Println(ele)
	}
	// Delete
	delete(set, "Foo")
	// Size
	size := len(set) // 0
	fmt.Println(size)
	// Exists
	_, exists := set["Foo"]
	fmt.Println(exists) // false
}
