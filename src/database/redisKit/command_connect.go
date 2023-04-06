package redisKit

import "context"

func (client *Client) Ping(ctx context.Context) (string, error) {
	return client.core.Ping(ctx).Result()
}
