package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/spf13/viper"
)

type (
	Config struct {
		A interface{} `json:"a"`
		B interface{} `json:"b"`
	}
)

func main() {
	v := viper.New()
	v.AddConfigPath("./")        // 设置读取路径：就是在此路径下搜索配置文件。
	v.SetConfigFile("test.json") // 设置被读取文件的全名，包括扩展名。
	//v.AllowEmptyEnv()
	err := v.ReadInConfig() // 读取配置文件： 这一步将配置文件变成了 Go语言的配置文件对象包含了 map，string 等对象。
	if err != nil {
		panic(err)
	}

	c := Config{}
	if err := v.Unmarshal(&c); err != nil {
		panic(err)
	}

	c.A = nil

	json, err := jsonKit.MarshalIndentToString(c, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(json)

	v.SafeWriteConfig()

	//fmt.Println(v.AllSettings())
	//
	///* 1 */
	//if err := v.ReadConfig(ioKit.NewReader([]byte(json))); err != nil {
	//	panic(err)
	//}
	//if err := v.WriteConfig(); err != nil {
	//	panic(err)
	//}

	/* 2 */
	if err := fileKit.WriteToFile([]byte(json), "./test.json"); err != nil {
		panic(err)
	}
}
