package main

import "C"
import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("a")    // 配置文件名称（没有扩展名）
	viper.AddConfigPath(".")    // 配置文件路径（可以设置多个）
	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {             // 处理读取配置文件的错误
		panic(err)
	}

	var m map[string]interface{}
	if err := viper.Unmarshal(&m); err != nil {
		panic(err)
	}
	str, _ := jsonKit.MarshalIndentToString(m, "", "    ")
	fmt.Println(str)
}
