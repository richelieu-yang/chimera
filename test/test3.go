package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/thirdParty/gaode/gaodeKit"
	"github.com/richelieu-yang/chimera/v2/src/thirdParty/gaode/weatherKit"
)

func main() {
	gaodeKit.MustSetUp("b15c36bf1df4c272e92f3f1875a127f1")

	live, err := weatherKit.GetLiveWeather("320205")
	if err != nil {
		panic(err)
	}
	fmt.Println(live)

	//fmt.Println(reqKit.Get("https://req.cool/zh/docs/prologue/quickstart/", nil))
}
