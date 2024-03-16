package main

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/component/database/nosql/mongodbKit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func main() {
	client, err := mongodbKit.NewClientSimply(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(nil)

	db := client.Database("db_a")
	collection := db.Collection("collection_b")
	{
		filter := bson.D{
			{Key: "name", Value: "Richelieu"},
		}
		update := bson.D{
			{Key: "$inc", Value: bson.D{
				{Key: "age", Value: 1},
			}},
		}
		opts := options.Update().SetComment("Test comment.")

		_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
		if err != nil {
			panic(err)
		}
	}
}
