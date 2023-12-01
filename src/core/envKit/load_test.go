package envKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	m, err := ReadFromFile("./_test.yaml", "./_test.env")
	if err != nil {
		panic(err)
	}
	json, err := jsonKit.MarshalIndentToString(m, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(json)
	if err := SetEnvs(m); err != nil {
		panic(err)
	}
	fmt.Println(GetEnv("c"))
}
