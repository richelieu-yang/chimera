package weatherKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/thirdParty/gaode/gaodeKit"
	"testing"
)

func TestWeather(t *testing.T) {
	gaodeKit.MustSetUp("b15c36bf1df4c272e92f3f1875a127f1")

	live, err := GetLive("320205")
	if err != nil {
		panic(err)
	}
	fmt.Println(live)

	todayCast, err := GetTodayCast("320205")
	if err != nil {
		panic(err)
	}
	fmt.Println(todayCast)
}
