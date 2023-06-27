package mapKit

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	type User struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	u := User{
		ID:   "666",
		Name: "测试员",
	}
	m := Encode(u)
	fmt.Println(m) // map[id:666 name:测试员]
	m1 := EncodeWithTag(&u, "json")
	fmt.Println(m1) // map[id:666 name:测试员]
}
