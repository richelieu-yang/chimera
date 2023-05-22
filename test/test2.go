package main

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gutil"
	jsoniter "github.com/json-iterator/go"
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
	var dest map[string]interface{}
	dest = DeepCopy(src).(map[string]interface{})

	fmt.Println(src)
	fmt.Println(dest)

	src["b"] = true
	b.Id = 777

	fmt.Println(jsonKit.MarshalToStringWithJsoniterApi(jsoniter.ConfigCompatibleWithStandardLibrary, src))
	fmt.Println(jsonKit.MarshalToStringWithJsoniterApi(jsoniter.ConfigCompatibleWithStandardLibrary, dest))
}

func DeepCopy(src interface{}) interface{} {
	return gutil.Copy(src)
}
