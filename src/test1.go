package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Clean("./1.txt"))      // "1.txt"
	fmt.Println(filepath.Clean("/root/.././c")) // "/c"
}
