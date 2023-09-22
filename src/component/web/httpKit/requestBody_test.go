package httpKit

import (
	"fmt"
	"testing"
)

func TestToRequestBodyString(t *testing.T) {
	m := map[string][]string{
		"a": []string{"test"},
		"b": []string{"测试"},
	}
	fmt.Println(ToRequestBodyString(m)) // a=test&b=%E6%B5%8B%E8%AF%95
}
