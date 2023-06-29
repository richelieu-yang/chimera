package mapKit

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func TestEncode(t *testing.T) {
	type User struct {
		gorm.Model

		Name string `json:"name"`
	}

	u := &User{
		Model: gorm.Model{
			ID: 666,
		},
		Name: "测试员",
	}

	/* 会转换为map[string]interface{}嵌套map[string]interface{} */
	m := Encode(u)
	// map[Model:map[CreatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:map[Time:0001-01-01 00:00:00 +0000 UTC Valid:false] ID:666 UpdatedAt:0001-01-01 00:00:00 +0000 UTC] name:测试员]
	fmt.Println(m)
	fmt.Println(len(m)) // 2
	m1 := EncodeWithTag(&u, "json")
	// map[Model:map[CreatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:map[Time:0001-01-01 00:00:00 +0000 UTC Valid:false] ID:666 UpdatedAt:0001-01-01 00:00:00 +0000 UTC] name:测试员]
	fmt.Println(m1)
	fmt.Println(len(m1)) // 2
}
