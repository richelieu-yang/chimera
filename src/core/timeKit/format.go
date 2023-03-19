package timeKit

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/sliceKit"
	"time"
)

// FormatCurrentTime 格式化当前时间为字符串
/*
e.g.
() 		=> "2022-08-13 14:54:44.336"
("") 	=> ""
*/
func FormatCurrentTime(formats ...TimeFormat) string {
	format := sliceKit.GetFirstItemWithDefault(FormatCommon, formats...)

	return FormatTimeToString(time.Now(), format)
}

// FormatTimeToString 格式化时间为字符串
/*
@param t		不用担心t为nil的情况，详见下面的说明
@param formats 	不传的话用默认值；传多个（包括一个）的话用第一个

一个方法如果接受类型为time.Time的参数，那么不用考虑该参数为nil的情况，因为：
（1）time.Time类型变量的零值不为nil；
（2）调用时，该参数位置不能直接传参nil（IDEA报错：Cannot use 'nil' as the type time.Time）；
（3）time.Time类型变量不能被赋值为nil（IDEA报错：Cannot use 'nil' as the type time.Time）。
*/
func FormatTimeToString(t time.Time, format TimeFormat) string {
	return t.Format(string(format))
}

// FormatDurationToString
/*
PS:
(1) %v 和 %s 都可以；
(2) strKit.ToString 也支持.

e.g.
(time.Minute*63 + time.Second*15) => "1h3m15s"
*/
func FormatDurationToString(t time.Duration) string {
	return fmt.Sprintf("%s", t)
}
