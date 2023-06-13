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

// NewDailyWriteCloser
/*
@param options 可选配置: WithCompress() || WithMaxAge() || WithMaxBackups()
*/
func NewDailyWriteCloser(filePath string, options ...LumberjackOption) (io.WriteCloser, error) {
	options = append(options, WithMaxSize(math.MaxInt64))
	wc, err := NewLumberjackWriteCloser(filePath, options...)
	if err != nil {
		return nil, err
	}

	c, _, err := cronKit.NewCronWithTask("0 0 0 * * *", func() {
		_, _ = wc.Write([]byte("rotate by c"))
		if err := wc.Rotate(); err != nil {
			logrus.WithError(err).Error("fail to rotate")
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
