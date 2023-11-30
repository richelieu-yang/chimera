package carbonKit

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"testing"
)

func TestJson(t *testing.T) {
	type Person struct {
		Name     string
		Birthday carbon.DateTime
	}

	p := Person{
		Name: "张三",
		Birthday: carbon.DateTime{
			Carbon: carbon.Parse("2022-08-08 12:12:12"),
		},
	}
	fmt.Println(jsonKit.MarshalIndentToString(p, "", "    "))
	/*
		{
		    "Name": "张三",
		    "Birthday": "2022-08-08 12:12:12"
		} <nil>
	*/
}
