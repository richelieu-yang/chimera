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
func SetSystemTime(t time.Time, rootPassword string) error {
	// 将时间转换为（date命令认可的）字符串
	format := "010215042006.05"
	timeStr := FormatTimeToString(t, TimeFormat(format))

	var script string
	if strKit.IsEmpty(rootPassword) {
		script = strKit.Format("date %s", timeStr)
	} else {
		script = strKit.Format(`echo "%s" | sudo -S date %s`, rootPassword, timeStr)
	}

	_, err := cmdKit.ExecuteToString("sh", "-c", script)
	return err
}

// CorrectSystemTime （根据网络时间）纠正系统时间
func CorrectSystemTime(rootPassword string) (t time.Time, err error) {
	t, _, err = GetNetworkTime()
	if err != nil {
		return
	}
	err = SetSystemTime(t, rootPassword)
	return
}
