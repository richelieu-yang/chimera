package validateKit

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	type inner struct {
		Sex string `validate:"required"`
	}
	type bean struct {
		Id    int    `validate:"required"`
		Name  string `validate:"required"`
		Inner *inner
	}

	b := &bean{
		Id:   1,
		Name: "",
		Inner: &inner{
			Sex: "",
		},
	}

	if err := Struct(b); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}

	if err := StructExcept(b, "Name", "Inner.Sex"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
