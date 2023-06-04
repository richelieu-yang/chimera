package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Create("a.log"))
}
