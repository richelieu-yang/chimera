//go:build !windows

package timeKit

import (
	"github.com/richelieu42/go-scales/src/cmdKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"time"
)

// SetSystemTime 设置系统时间（机器时间）
/*
PS:
(1) 通过date命令设置root权限，需要root权限.
(2) 只能精确到"秒".

@param password root用户的密码
*/
func SetSystemTime(t time.Time, password string) error {
	// 将时间转换为（date命令认可的）字符串
	format := "010215042006.05"
	timeStr := FormatTimeToString(t, TimeFormat(format))

	var script string
	if strKit.IsEmpty(password) {
		script = strKit.Format("date %s", timeStr)
	} else {
		script = strKit.Format(`echo "%s" | sudo -S date %s`, password, timeStr)
	}

	_, err := cmdKit.ExecuteToString("sh", "-c", script)
	return err
}
