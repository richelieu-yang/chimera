package wireKit

import "github.com/google/wire"

var (
	NewSet func(...interface{}) wire.ProviderSet = wire.NewSet
)
