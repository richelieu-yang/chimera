package main

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/yaml/k8sYamlKit"
)

type Person struct {
	Name string `json:"name"` // Affects YAML field names too.
	Age  int    `json:"age"`
}

func main() {
	path := "test3.yaml"

	data, err := fileKit.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var p Person
	if err := k8sYamlKit.Unmarshal(data, &p); err != nil {
		panic(err)
	}
	p.Age++

	data, err = k8sYamlKit.Marshal(&p)
	if err != nil {
		panic(err)
	}

	if err := fileKit.WriteToFile(data, path, 0777); err != nil {
		panic(err)
	}
}
