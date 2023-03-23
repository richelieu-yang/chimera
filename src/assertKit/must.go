package assertKit

import (
	"github.com/richelieu42/chimera/src/funcKit"
	"github.com/sirupsen/logrus"
)

// Must
/*
参考: logx.Must().

@param err 如果不为nil，进程将退出(os.Exit(1))
*/
func Must(err error) {
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caller": funcKit.GetCaller(2),
		}).Fatal(err)
	}
}
