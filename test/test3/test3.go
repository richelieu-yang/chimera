package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonRespKit"
)

func main() {
	fmt.Println(jsonRespKit.Seal("0", nil))
}
