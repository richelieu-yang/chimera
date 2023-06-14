package ioKit

import (
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"io"
	"math"
)

type DailyWriteCloser struct {
	writeCloser io.WriteCloser
	cron        *cron.Cron
}

func (dwc *DailyWriteCloser) Write(p []byte) (int, error) {
	return dwc.writeCloser.Write(p)
}

func (dwc *DailyWriteCloser) Close() error {
	return dwc.writeCloser.Close()
}

// NewDailyWriteCloser 每天凌晨0点rotate1次.
/*
@param options 可选配置，参考 NewRotatableWriteCloser()
*/
func NewDailyWriteCloser(filePath string, options ...LumberjackOption) (io.WriteCloser, error) {
	wc, err := NewRotatableWriteCloser(filePath, math.MaxInt64, options...)
	if err != nil {
		return nil, err
	}

	c, _, err := cronKit.NewCronWithTask("0 0 0 * * *", func() {
		_, _ = wc.Write([]byte("rotate by cron"))
		if err := wc.Rotate(); err != nil {
			text := "fail to rotate by cron, error: " + err.Error()
			_, _ = wc.Write([]byte(text))
			logrus.Error(text)
		}
	})
	if err != nil {
		return nil, err
	}
	c.Start()

	return &DailyWriteCloser{
		writeCloser: wc,
		cron:        c,
	}, nil
}
