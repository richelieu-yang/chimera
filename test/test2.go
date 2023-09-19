package main

import (
	"fmt"
	"os"
	"sigs.k8s.io/yaml"
)

func main() {
	// 读取yaml文件
	file, err := os.Open("example.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data interface{}
	err = yaml.NewDecoder(file).Decode(&data)
	if err != nil {
		panic(err)
	}

	// 将数据转换为yaml字符串
	yamlStr, err := yaml.Marshal(data)
	if err != nil {
		panic(err)
	}

	// 将yaml字符串转换为map
	var result map[string]interface{}
	err = yaml.Unmarshal(yamlStr, &result)
	if err != nil {
		panic(err)
	}

	// 输出转换后的map
	fmt.Println(result)

	// 将map转换为yaml文件
	file, err = os.Create("example.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = yaml.NewEncoder(file).Encode(result)
	if err != nil {
		panic(err)
	}
}
