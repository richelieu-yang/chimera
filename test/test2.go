package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/core/timeKit"
	"time"
)

func main() {
	fmt.Println(timeKit.ToRelativeString(time.Now().Add(time.Hour * 16)))
	fmt.Println(timeKit.ToRelativeString(time.Now().Add(-time.Hour * 16)))
	fmt.Println(timeKit.ToRelativeString(time.Now().Add(time.Hour * 24 * 21)))
	fmt.Println(timeKit.ToRelativeString(time.Now().Add(time.Hour*24*21 + time.Second)))

}
