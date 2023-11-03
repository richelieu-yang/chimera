package pushKit

import (
	"github.com/panjf2000/ants/v2"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
	"github.com/sirupsen/logrus"
)

var pool *ants.Pool

func MustSetUp(p *ants.Pool) {
	if err := Setup(p); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

// Setup
/*
@param p 需要自行决定: cap大小、是否自定义输出...
*/
func Setup(p *ants.Pool) error {
	if p.IsClosed() {
		return errorKit.New("pool has already been closed")
	}
	capacity := p.Cap()
	if capacity > 0 {
		tag := "gte=2000"
		if err := validateKit.Var(capacity, tag); err != nil {
			return errorKit.Wrap(err, "Capacity(%d) of pool is invalid(tag: %s) when it's greater than zero", capacity, tag)
		}
	}

	pool = p
	return nil
}

func isAvailable() error {
	if pool == nil {
		return NotSetupError
	}
	return nil
}
