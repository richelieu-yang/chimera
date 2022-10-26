package wsKit

import "github.com/richelieu42/go-scales/src/core/errorKit"

var (
	NotInitializedError = errorKit.Simple("Connection hasn't been initialized correctly")

	DisposedError = errorKit.Simple("Connection has already been disposed")

	NoConnError = errorKit.Simple("no conn")
)
