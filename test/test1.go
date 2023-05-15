package main

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
)

func main() {
	t := mimetype.Detect(nil)
	fmt.Println(t)

	t.
}
