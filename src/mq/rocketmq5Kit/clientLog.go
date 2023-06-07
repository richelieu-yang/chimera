package rocketmq5Kit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/osKit"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"os"
)

func initClientLog(path string) error {
	if strKit.IsBlank(path) {
		// (1) 输出到控制台
		if err := os.Setenv(rmq_client.ENABLE_CONSOLE_APPENDER, "true"); err != nil {
			return err
		}
	} else {
		// (2) 输出到日志文件
		if err := fileKit.AssertNotExistOrIsFile(path); err != nil {
			return err
		}
		dir := pathKit.GetParentDir(path)
		if err := fileKit.MkDirs(dir); err != nil {
			return err
		}
		name := fileKit.GetBaseName(path)

		if err := osKit.SetEnvs(map[string]string{
			rmq_client.ENABLE_CONSOLE_APPENDER: "false",
			rmq_client.CLIENT_LOG_ROOT:         dir,
			rmq_client.CLIENT_LOG_FILENAME:     name,
		}); err != nil {
			return err
		}
	}

	rmq_client.ResetLogger()
	return nil
}
