package timeKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"time"
)

// ConvertToLocalLocation 转换为本地时区的时间.
func ConvertToLocalLocation(t time.Time) time.Time {
	return t.In(time.Local)
}

func ConvertToUTC(t time.Time) time.Time {
	return t.In(time.UTC)
}

// ConvertByLocation 时区转换
/*
@param loc 目标时区（不能为nil！！！）

e.g. UTC+8 转 UTC+0
2022-05-05 14:33:40.562899 +0800 CST m=+0.001585418 => 2022-05-05 06:33:40.562899 +0000 UTC
*/
func ConvertByLocation(t time.Time, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return time.Time{}, errorKit.Simple("loc is nil")
	}
	return t.In(loc), nil
}

// LoadLocation
/*
LoadLocation的输入参数的取值，除了该函数的源代码中可看到的”UTC”、”Local”，其余的值其实是遵照“IANA Time Zone”的规则，可以解压$GOROOT/lib/time/zoneinfo.zip 这个文件打开查看。
在Asia这个目录，我看到了Chongqing，Hong_Kong，但没Beijing。在国外获取中国北京时间，要用”PRC”，当然”Asia/Chongqing”也是个方法
参考：https://blog.csdn.net/qq_26981997/article/details/53454606

@param name e.g. "Asia/Chongqing"
*/
func LoadLocation(name string) (*time.Location, error) {
	return time.LoadLocation(name)
}
