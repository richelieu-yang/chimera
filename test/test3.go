package main

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/component/database/nosql/mongodbKit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
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
		filter := bson.D{{"name", "张3"}}
		findOptions := options.Find()
		// 默认0（不限制）
		findOptions.SetLimit(2)

		cur, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			panic(err)
		}
		// 关闭游标
		defer cur.Close(context.TODO())

		// 遍历游标，解码文档
		var results []*Student
		for cur.Next(context.TODO()) {
			elem := &Student{}
			err := cur.Decode(elem)
			if err != nil {
				panic(err)
			}
			results = append(results, elem)
		}
		if err := cur.Err(); err != nil {
			panic(err)
		}

		// 如果不存在满足条件的文档，此处 len(results) == 0 （上面并不会返回error）

		for _, s := range results {
			fmt.Printf("name: %s, age: %d\n", s.Name, s.Age)
		}
	}
}
