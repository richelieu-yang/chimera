package jsonKit

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func TestMarshalToString(t *testing.T) {
	type user struct {
		gorm.Model

		Name string `json:"name"`
	}

	u := &user{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "test",
	}
	fmt.Println(MarshalToString(u, WithIndent("    ")))
}
