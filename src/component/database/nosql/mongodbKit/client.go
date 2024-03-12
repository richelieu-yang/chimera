package mongodbKit

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// NewClient
/*
@param uri e.g."mongodb://localhost:27017"
*/
func NewClient(ctx context.Context, uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	/* ping */
	ctx1, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := client.Ping(ctx1, readpref.Primary()); err != nil {
		_ = client.Disconnect(context.TODO())
		return nil, err
	}

	return client, nil
}
