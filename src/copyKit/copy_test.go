package copyKit

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	type bean struct {
		Id int
	}

	b := &bean{
		Id: 666,
	}
	src := map[string]interface{}{
		"b":   false,
		"tmp": b,
	}

	dest, err := DeepCopy(src)
	if err != nil {
		panic(err)
	}

	fmt.Println(src)
	fmt.Println(dest)

	src["b"] = true
	b.Id = 777

	fmt.Println(jsonKit.MarshalToStringWithAPI(jsoniter.ConfigCompatibleWithStandardLibrary, src))
	fmt.Println(jsonKit.MarshalToStringWithAPI(jsoniter.ConfigCompatibleWithStandardLibrary, dest))
}
