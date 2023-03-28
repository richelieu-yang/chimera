package confKit

import (
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/conf"
)

// MustLoad 读取配置文件并反序列化为 指定结构体指针ptr
/*
@param opts e.g. conf.UseEnv()
*/
func MustLoad(path string, ptr any, opts ...conf.Option) {
	if err := load(path, ptr, opts...); err != nil {
		logrus.Fatal(err)
	}
}

func load(path string, ptr any, opts ...conf.Option) error {
	return conf.Load(path, ptr, opts...)
}
