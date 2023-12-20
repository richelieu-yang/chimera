package oauth2Kit

import (
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/redis/go-redis/v9"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/reflectKit"
)

//var (
//	NewRedisStore func(opts *redis.Options, keyNamespace ...string) *oredis.TokenStore = oredis.NewRedisStore
//)

func NewRedisStore(client redis.UniversalClient) (*oredis.TokenStore, error) {
	// Richelieu: 由于 go-oauth2/redis 使用的是 go-redis/redis/v8，版本太低了，所以此处使用反射修改未导出字段.
	store := &oredis.TokenStore{}
	if err := reflectKit.SetField(store, "cli", client); err != nil {
		return nil, errorKit.Wrap(err, "Fail to set unexported field 'cli'")
	}
	return store, nil
}
