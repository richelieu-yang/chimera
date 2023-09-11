package redisKit

import "context"

func (client *Client) Ping(ctx context.Context) (string, error) {
	return client.universalClient.Ping(ctx).Result()
}
