package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/web/httpClientKit"
)

func main() {
	fmt.Println(httpClientKit.Post("https://www.moulem.com/?fr=au12"))
	fmt.Println(httpClientKit.Upload("https://www.moulem.com/?fr=au12", nil))
}
