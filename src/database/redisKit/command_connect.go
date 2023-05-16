package redisKit

import "context"

func (client *Client) Ping(ctx context.Context) (string, error) {
	cmd := client.universalClient.Ping(ctx)
	return cmd.Result()
}
