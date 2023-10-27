package timeKit

import "time"

// ToDefaultIfInvalid 如果t的值无效，将返回默认值def；否则返回t.
func ToDefaultIfInvalid(t, def time.Duration) time.Duration {
	if t <= 0 {
		return def
	}
	return t
}
