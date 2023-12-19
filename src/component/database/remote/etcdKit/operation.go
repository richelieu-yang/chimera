package etcdKit

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// Exist key是否存在？
func Exist(kv clientv3.KV, ctx context.Context, key string) (bool, error) {
	getResp, err := kv.Get(ctx, key)
	if err != nil {
		return false, err
	}
	return getResp.Kvs != nil, nil
}
