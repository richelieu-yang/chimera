package main

import (
	"github.com/richelieu42/chimera/v2/src/copyKit"
)

func main() {
	m := map[string]interface{}{
		"a": 0,
	}
	err := copyKit.Copy(&m, nil)
	if err != nil {
		panic(err)
	}
}
