package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	// 读取 YAML 文件
	file, err := os.Open("example.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 解析 YAML 文件
	var data map[string]interface{}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	// 修改数据
	data["name"] = "new-name"

	// 将数据转换为 YAML 字符串
	yamlBytes, err := yaml.Marshal(data)
	if err != nil {
		panic(err)
	}

	// 将 YAML 字符串写入文件
	file, err = os.Create("example.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(yamlBytes)
	if err != nil {
		panic(err)
	}

	fmt.Println("文件已保存")
}
