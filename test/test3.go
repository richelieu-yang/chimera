package main

import "fmt"

func main() {
	for i := 0; i < 26; i++ {
		fmt.Println(string('a'+i), string('a'+25-i))
		fmt.Println("---")
	}
}
