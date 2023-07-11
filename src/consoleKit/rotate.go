package consoleKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"time"
)

// RotateOutput
/*
@param output 	控制台输出文件（e.g.nohup.out）的路径
@param backDir	备份文件存放的目录
*/
func RotateOutput(output, backupDir, spec string) (*cron.Cron, error) {
	c, _, err := cronKit.NewCronWithTask(spec, func() {
		if err := rotate(output, backupDir); err != nil {
			logrus.Errorf("%+v", err)
		}
	})
	if err != nil {
		return nil, err
	}
	c.Start()
	return c, nil
}

func rotate(output, backupDir string) error {
	if err := fileKit.AssertExistAndIsFile(output); err != nil {
		return err
	}
	if err := fileKit.MkDirs(backupDir); err != nil {
		return err
	}

	dateStr := timeKit.FormatTimeToString(time.Now().Add(-timeKit.Day), "2006-01-02")
	name := fileKit.GetName(output)
	ext := fileKit.GetExt(output)
	target := pathKit.Join(backupDir, fmt.Sprintf("%s-%s%s", name, dateStr, ext))
	// 复制文件
	if err := fileKit.CopyFile(output, target); err != nil {
		return err
	}
	// 清空文件
	if err := fileKit.Truncate(output, 0); err != nil {
		return err
	}
	return nil
}
