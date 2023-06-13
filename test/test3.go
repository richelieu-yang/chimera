package main

import (
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
)

func main() {
	logrusKit.MustSetUp(nil)

	cron, _, err := cronKit.NewCron("30 * * * * *", func() {

	})
	if err != nil {
		panic(err)
	}
	cron.Start()
}
