package jsonKit

import (
	"fmt"
	"testing"
)

func TestUnmarshalFromString(t *testing.T) {
	json := `{"a":"b"}`
	var m map[string]any = nil

	if err := UnmarshalFromString(json, &m); err != nil {
		panic(err)
	}
	fmt.Println(m["a"]) // "b"
}
