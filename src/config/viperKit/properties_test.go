package viperKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"testing"
)

// TestLoadProperties 读取"_test.properties".
func TestLoadProperties(t *testing.T) {
	type redis struct {
		Cluster []string `json:"cluster"`
	}
	type config struct {
		Redis redis `json:"redis"`
	}

	path := "_test.properties"
	c := &config{}
	_, err := ReadFileAs(path, nil, c)
	if err != nil {
		panic(err)
	}

	str, _ := jsonKit.MarshalIndentToString(c, "", "    ")
	fmt.Println(str)
}
