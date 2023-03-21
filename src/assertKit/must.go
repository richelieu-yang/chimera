package assertKit

import (
	"github.com/richelieu42/chimera/src/funcKit"
	"log"
)

// Must
/*
参考: logx.Must().

@param err 如果不为nil，进程将退出(os.Exit(1))
*/
func Must(err error) {
	if err != nil {
		log.Fatalf("[FATAL] caller(%s) and error:\n%+v", funcKit.GetCaller(2), err)

		//logrus.WithFields(logrus.Fields{
		//	"caller": funcKit.GetCaller(2),
		//}).Fatalf("%+v", err)
	}
}
