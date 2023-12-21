package typeKit

import (
	"fmt"
	"testing"
)

func TestGetTypeString(t *testing.T) {
	fmt.Println(GetTypeString(0)) // int
	fmt.Println(GetTypeString(1)) // int

	fmt.Println(GetTypeString("")) // string

	var obj interface{}
	fmt.Println(GetTypeString(obj)) // <nil>

	type bean struct {
	}
	var b *bean = nil
	fmt.Println(GetTypeString(b)) // *typeKit.bean
	var obj1 interface{} = b
	fmt.Println(GetTypeString(obj1)) // *typeKit.bean
}
