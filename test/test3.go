package main

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/component/database/nosql/mongodbKit"
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

	// 引用名为"myNewDatabase"的数据库
	db := client.Database("1")

	// 引用或创建名为"users"的集合
	collection := db.Collection("users")

	// 创建一个用户文档
	user := User{Name: "John Doe", Email: "john.doe@example.com"}

	// 插入文档到集合中，这将自动创建数据库和集合（如果它们不存在）
	if _, err := collection.InsertOne(context.Background(), user); err != nil {
		panic(err)
	}

	fmt.Println("ok")
}
