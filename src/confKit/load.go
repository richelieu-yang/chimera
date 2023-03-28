package confKit

import (
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/conf"
)

// MustLoad 读取配置文件并反序列化为 指定结构体指针ptr
/*
!!!: 结构体可以参考 go-zero/rest/config.go中的RestConf，可以通过tag控制配置的值（默认值default、范围range、可选optional...）

@param ptr	（不能为nil）结构体实例的指针
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
