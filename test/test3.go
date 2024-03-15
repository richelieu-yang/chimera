package main

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/component/database/nosql/mongodbKit"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	client, err := mongodbKit.NewClientSimply(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(nil)

	filter := bson.M{}
	names, err := client.ListDatabaseNames(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(names)      // [admin config db_a local]
	fmt.Println(len(names)) // 4
}
