package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/cmdKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/core/timeKit"
	"time"
)

func main() {
	t, err := timeKit.ParseStringToTime(string(timeKit.CommonFormat), "2006-01-02 15:04:05.000")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	if err := SetSystemTime(t, "Cyy7587141200"); err != nil {
		panic(err)
	}

	//result, err := cmdKit.ExecuteToString("sh", "-c", `echo "Cyy7587141200" | sudo -S date 010203042010.05`)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(result)
}

// SetSystemTime 设置系统时间（机器时间）
/*
PS:
通过date命令设置root权限，需要root权限.

@param password root用户的密码
*/
func SetSystemTime(t time.Time, password string) error {
	format := "010215042006.05"
	timeStr := timeKit.FormatTimeToString(t, timeKit.TimeFormat(format))

	var script string
	if strKit.IsEmpty(password) {
		script = strKit.Format("date %s", timeStr)
	} else {
		script = strKit.Format(`echo "%s" | sudo -S date %s`, password, timeStr)
	}

	_, err := cmdKit.ExecuteToString("sh", "-c", script)
	return err
}
