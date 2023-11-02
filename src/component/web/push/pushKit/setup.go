package pushKit

import (
	"github.com/panjf2000/ants/v2"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

var pool *ants.Pool

// Setup
/*
@param p
*/
func Setup(p *ants.Pool) error {
	if p.IsClosed() {
		return errorKit.New("pool has already been closed")
	}
	capacity := p.Cap()
	if capacity > 0 {
		tag := "gte=2000"
		if err := validateKit.Var(capacity, tag); err != nil {
			return errorKit.Wrap(err, "Capacity(%d) of pool is invalid(%s) when it's greater than zero", capacity, tag)
		}
	}

	pool = p
	return nil
}
