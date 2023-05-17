package randomKit

import "github.com/richelieu42/chimera/v2/src/core/errorKit"

// Int
/*
@return 范围: [min, max)
*/
func Int(min, max int) (int, error) {
	if min >= max {
		return 0, errorKit.Simple("min(%d) is greater than or equal to max(%d)", min, max)
	}

	i := r.Intn(max - min)
	return i + min, nil
}
