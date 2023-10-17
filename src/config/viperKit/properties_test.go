package viperKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"testing"
)

// TestLoadProperties 读取"_test.properties".
func TestLoadProperties(t *testing.T) {
	var m map[string]interface{}

	//viper.SetConfigName("_test") // 配置文件名称（没有扩展名）
	//viper.AddConfigPath(".")     // 配置文件路径（可以设置多个）
	//err := viper.ReadInConfig()  // 读取配置数据
	//if err != nil {              // 处理读取配置文件的错误
	//	panic(err)
	//}
	//if err := viper.Unmarshal(&m); err != nil {
	//	panic(err)
	//}

	path := "_test.properties"
	_, err := ReadFileAs(path, nil, &m)
	if err != nil {
		panic(err)
	}

	str, _ := jsonKit.MarshalIndentToString(m, "", "    ")
	fmt.Println(str)
}
