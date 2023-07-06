package sonicKit

import (
	"fmt"
	"github.com/bytedance/sonic"
	"testing"
)

func TestMarshalToStringByAPIWithIndent(t *testing.T) {
	m := map[string]interface{}{
		"a": 1,
		"b": 2,
		"":  3,
	}
	str, err := MarshalToStringByAPIWithIndent(sonic.ConfigDefault, m, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
