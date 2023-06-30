package strKit

import "strings"

var (
	// Trim
	/*

	 */
	Trim func(s, cutset string) string = strings.Trim

	// TrimLeft
	/*

	 */
	TrimLeft func(s, cutset string) string = strings.TrimLeft

	// TrimRight
	/*

	 */
	TrimRight func(s, cutset string) string = strings.TrimRight
)
