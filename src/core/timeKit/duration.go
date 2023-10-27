package timeKit

import "time"

// ToDefaultDurationIfInvalid 如果d的值无效，将返回默认值def；否则返回d.
func ToDefaultDurationIfInvalid(d, def time.Duration) time.Duration {
	if d <= 0 {
		return def
	}
	return d
}
