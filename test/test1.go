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
	dest, err := copyKit.DeepCopy(src)
	if err != nil {
		panic(err)
	}

	// {"b":false,"tmp":{"Id":666}} <nil>
	fmt.Println(jsonKit.MarshalToString(src, jsonKit.WithApi(jsoniter.ConfigCompatibleWithStandardLibrary)))
	// {"b":false,"tmp":{"Id":666}} <nil>
	fmt.Println(jsonKit.MarshalToString(dest, jsonKit.WithApi(jsoniter.ConfigCompatibleWithStandardLibrary)))

	src["b"] = true
	b.Id = 777

	// {"b":true,"tmp":{"Id":777}} <nil>
	fmt.Println(jsonKit.MarshalToString(src, jsonKit.WithApi(jsoniter.ConfigCompatibleWithStandardLibrary)))
	// {"b":false,"tmp":{"Id":666}} <nil>
	fmt.Println(jsonKit.MarshalToString(dest, jsonKit.WithApi(jsoniter.ConfigCompatibleWithStandardLibrary)))
}
