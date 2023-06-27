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
	fmt.Println(m) // map[ID:666 Name:测试员]
	m1 := Encode(&u)
	fmt.Println(m1) // map[ID:666 Name:测试员]

	Encode(nil)
}
