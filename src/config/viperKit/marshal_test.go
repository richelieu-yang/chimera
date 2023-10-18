package viperKit

import "testing"

func TestMarshalToFile(t *testing.T) {
	type config struct {
		A string
		B []string
	}

	c := config{
		A: "a",
		B: []string{"0", "1", "2"},
	}
	if err := MarshalToFile(c, "_test.yaml"); err != nil {
		panic(err)
	}
	if err := MarshalToFile(c, "_test.properties"); err != nil {
		panic(err)
	}
}
