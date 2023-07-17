package cronKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/robfig/cron/v3"
)

// NewCron
/*
!!!: 想要通过修改机器时间来验证的话，需要先改时间，再启动cron.

定时任务-表达式
	https://goframe.org/pages/viewpage.action?pageId=30736411
Go 每日一库之定时任务库：cron
	https://mp.weixin.qq.com/s/swdijAro2k8LuYu7q_La1A
cron表达式，每天凌晨0点执行定时任务
	https://www.cnblogs.com/yddwinter/p/16033633.html


e.g. spec
"0 0 0 * * *"			每天凌晨0点执行
"* * * * * *"			每秒执行
"30 * * * * *"			每分钟的第30s，执行一次
"15,30 * * * * *"		每分钟的第15s、第30s，各执行一次
"@every 10s"			从执行Run() || Start()开始，每 10s	 	执行一次
"@every 1m"				从执行Run() || Start()开始，每 1min	执行一次
"@hourly"				从执行Run() || Start()开始，每 1h 		执行一次
"@every 1h30m"			从执行Run() || Start()开始，每 1.5h 	执行一次
*/
func NewCron() *cron.Cron {
	// cron.WithSeconds(): 带"秒"
	return cron.New(cron.WithSeconds())
}

// NewCronWithTask
/*
PS:
(1) 第1个返回值: 可以调用 Run() 或 Start() 以启动.（Run()会阻塞调用此方法的goroutine；Start()不会阻塞）

@param spec "@every 10s" || "@every 1m"，更多可参考"Golang - 1.docx"
*/
func NewCronWithTask(spec string, task func()) (*cron.Cron, cron.EntryID, error) {
	if strKit.IsBlank(spec) {
		return nil, 0, errorKit.New("spec is blank")
	}
	if task == nil {
		return nil, 0, errorKit.New("task == nil")
	}

	c := NewCron()
	entryId, err := c.AddFunc(spec, task)
	if err != nil {
		return nil, 0, err
	}
	return c, entryId, nil
}

func NewCronWithJob(spec string, job cron.Job) (*cron.Cron, cron.EntryID, error) {
	if strKit.IsBlank(spec) {
		return nil, 0, errorKit.New("spec is blank")
	}
	if job == nil {
		return nil, 0, errorKit.New("job == nil")
	}

	c := NewCron()
	entryId, err := c.AddJob(spec, job)
	if err != nil {
		return nil, 0, err
	}
	return c, entryId, nil
}
