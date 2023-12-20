package mapKit

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	m0 := map[string]interface{}{
		"a": 1,
	}
	m1 := map[string]interface{}{
		"m0": m0,
		"b":  2,
	}

	dolly := Copy(m1)
	m0["a"] = "a"
	m1["b"] = "b"

	fmt.Println(m1)    // map[b:b m0:map[a:a]]
	fmt.Println(dolly) // map[b:2 m0:map[a:a]]
}

func TestDeepCopy(t *testing.T) {
	m0 := map[string]interface{}{
		"a": 1,
	}
	m1 := map[string]interface{}{
		"m0": m0,
		"b":  2,
	}

	dolly := DeepCopy(m1)
	m0["a"] = "a"
	m1["b"] = "b"

	fmt.Println(m1)    // map[b:b m0:map[a:a]]
	fmt.Println(dolly) // map[b:2 m0:map[a:1]]
}
