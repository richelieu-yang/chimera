package pushKit

import (
	"github.com/panjf2000/ants/v2"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
	"github.com/sirupsen/logrus"
)

var pool *ants.Pool

func MustSetUp(antPool *ants.Pool, logrusLogger *logrus.Logger) {
	if err := Setup(antPool, logrusLogger); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

// Setup
/*
@param antPool	需要自行决定: cap大小、是否自定义输出...
@param logger 	可以为nil
*/
func Setup(antPool *ants.Pool, logrusLogger *logrus.Logger) error {
	/* pool */
	if antPool.IsClosed() {
		return errorKit.New("pool has already been closed")
	}
	capacity := antPool.Cap()
	if capacity > 0 {
		tag := "gte=2000"
		if err := validateKit.Var(capacity, tag); err != nil {
			return errorKit.Wrap(err, "Capacity(%d) of pool is invalid(tag: %s) when it's greater than zero", capacity, tag)
		}
	}
	pool = antPool

	/* logger */
	if logrusLogger != nil {
		if err := SetDefaultLogger(logrusLogger); err != nil {
			return err
		}
	}

	return nil
}

func isAvailable() error {
	if pool == nil {
		return NotSetupError
	}
	return nil
}
