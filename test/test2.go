package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/component/web/reqKit"
)

func main() {
	fmt.Println(reqKit.Post("https://www.moulem.com/", nil, nil))
}
