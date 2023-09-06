package jsonKit

import "testing"

func TestMarshalToFile(t *testing.T) {
	m := map[string]interface{}{
		"a": 1,
		"b": []string{"0", "1", "2"},
	}
	if err := MarshalToFile(m, "_test.json"); err != nil {
		panic(err)
	}
	if err := MarshalToFileWithAPI(GetStdApi(), m, "_test1.json"); err != nil {
		panic(err)
	}
}
