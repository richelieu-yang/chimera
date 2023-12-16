package redisKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

var (
	NotSetupError = errorKit.New("Haven’t been set up correctly")
)

// IsConsumerGroupNameAlreadyExistError
/*
TODO: 涉及 github.com/redis/go-redis/v9 源码， 后续看有没有好的解决方法.

PS: 与 XGroupCreateMkStream 搭配使用.
*/
func IsConsumerGroupNameAlreadyExistError(err error) bool {
	return strKit.ContainsIgnoreCase(err.Error(), "Consumer Group name already exists")
}
