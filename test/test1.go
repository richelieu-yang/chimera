package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/richelieu42/chimera/v2/src/copyKit"
	"github.com/richelieu42/chimera/v2/src/jsonKit"
)

type Bean struct {
	Id int
}

func main() {
	b := &Bean{
		Id: 666,
	}
	src := map[string]interface{}{
		"b":   false,
		"tmp": b,
	}
	dest := copyKit.DeepCopy(src).(map[string]interface{})

	fmt.Println(src)
	fmt.Println(dest)

	src["b"] = true
	b.Id = 777

	fmt.Println(jsonKit.MarshalToStringWithJsoniterApi(jsoniter.ConfigCompatibleWithStandardLibrary, src))
	fmt.Println(jsonKit.MarshalToStringWithJsoniterApi(jsoniter.ConfigCompatibleWithStandardLibrary, dest))
}
