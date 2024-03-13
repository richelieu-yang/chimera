package main

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/component/database/nosql/mongodbKit"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	type User struct {
		Name  string `bson:"name"`
		Email string `bson:"email"`
	}

	client, err := mongodbKit.NewClientSimply(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(nil)

	{
		db := client.Database("a")
		collection := db.Collection("b")

		// 创建一个用户文档
		user := &User{
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}
		// 插入文档到集合中，这将自动创建数据库和集合（如果它们不存在）
		rst, err := collection.InsertOne(context.Background(), user)
		if err != nil {
			panic(err)
		}

		id := rst.InsertedID.(primitive.ObjectID)
		fmt.Println(id.String()) // ObjectID("65f1720a39fddfbbd5d75315")
	}
}
