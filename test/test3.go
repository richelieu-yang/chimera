package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/confKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
)

type bean struct {
	Age int `json:"age"`
}

func main() {
	m := map[string]interface{}{
		"age": "11111",
	}
	jsonStr, err := jsonKit.MarshalToString(m)
	if err != nil {
		panic(err)
	}

	b := &bean{}
	//if err := confKit.LoadFromJsonText(jsonStr, b); err != nil {
	//	panic(err)
	//}

	if err := jsonKit.UnmarshalFromString(jsonStr, b); err != nil {
		panic(err)
	}

	if err := confKit.ReadAs([]byte(jsonStr), "json", nil, b); err != nil {
		panic(err)
	}

	fmt.Println(jsonKit.MarshalIndentToString(b, "", "    "))

	//path := "test3.yaml"
	//
	//data, err := fileKit.ReadFile(path)
	//if err != nil {
	//	panic(err)
	//}
	//
	//var p Person
	//if err := k8sYamlKit.Unmarshal(data, &p); err != nil {
	//	panic(err)
	//}
	//p.Age++
	//
	//data, err = k8sYamlKit.Marshal(&p)
	//if err != nil {
	//	panic(err)
	//}
	//
	//if err := fileKit.WriteToFile(data, path, 0777); err != nil {
	//	panic(err)
	//}
}
