package timeKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/robfig/cron/v3"
)

// NewCron
/*
@param spec 可参考"Golang - 1.docx"

@return 第1个返回值: 可以调用 Run() 或 Start() 以启动.（Run()会阻塞调用此方法的goroutine；Start()不会阻塞）
*/
func NewCron(spec string, task func()) (*cron.Cron, cron.EntryID, error) {
	if strKit.IsEmpty(spec) {
		return nil, 0, errorKit.Simple("spec is empty")
	}
	if task == nil {
		return nil, 0, errorKit.Simple("task == nil")
	}

	c := cron.New(cron.WithSeconds())
	entryId, err := c.AddFunc(spec, task)
	if err != nil {
		return nil, 0, err
	}
	return c, entryId, nil
}
