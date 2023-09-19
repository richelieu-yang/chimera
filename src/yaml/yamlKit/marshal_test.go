package yamlKit

import (
	"fmt"
	"testing"
)

func TestMarshalToString(t *testing.T) {
	{
		s := []string{"foo", "bar"}
		yamlStr, err := MarshalToString(s)
		if err != nil {
			panic(err)
		}
		fmt.Println(yamlStr)
	}
	fmt.Println("------")
	{
		m := map[string]interface{}{
			"a": 100,
			"b": 200,
		}
		yamlStr, err := MarshalToString(m)
		if err != nil {
			panic(err)
		}
		fmt.Println(yamlStr)
	}
}
