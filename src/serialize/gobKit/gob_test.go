package gobKit

import (
	"fmt"
	"testing"
)

func TestMarshalAndUnmarshal(t *testing.T) {
	m := map[interface{}]interface{}{
		"0": 3.1415926,
	}

	data, err := Marshal(m)
	if err != nil {
		panic(err)
	}

	var m1 map[interface{}]interface{}
	if err := Unmarshal(data, &m1); err != nil {
		panic(err)
	}
	fmt.Println(m1)
}
