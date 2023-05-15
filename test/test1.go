package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/web/mimeTypeKit"
)

func main() {
	mime := mimeTypeKit.Detect(nil)
	fmt.Println(mime.String()) // "text/plain"
}
