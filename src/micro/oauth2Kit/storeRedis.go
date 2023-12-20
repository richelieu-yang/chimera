package oauth2Kit

import (
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
)

var (
	NewRedisStore func(opts *redis.Options, keyNamespace ...string) *oredis.TokenStore = oredis.NewRedisStore
)
