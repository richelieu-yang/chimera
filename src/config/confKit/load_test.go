package confKit

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"testing"
	"time"
)

func TestMustLoad(t *testing.T) {
	type config1 struct {
		Name string `json:"name,default=test"`
	}

	type config struct {
		Time    time.Duration `json:"time"`
		Number  int           `json:"number,range=[1:100)"`
		Number1 int           `json:"number1,range=[1:)"`
		// 此时使用 "config1" 而非 "*config1"，避免nil的情况
		Config1 config1 `json:"config1"`
	}

	c := &config{}
	MustLoad("./_test.yaml", c)

	fmt.Printf("c.Time: %s\n", c.Time)
	fmt.Printf("c.Number: %d\n", c.Number)
	fmt.Printf("c.Number1: %d\n", c.Number1)
	fmt.Printf("c.Config1.Name: %s\n", c.Config1.Name)
}

func TestMustLoad1(t *testing.T) {
	type args struct {
		path    string
		ptr     any
		options []conf.Option
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MustLoad(tt.args.path, tt.args.ptr, tt.args.options...)
		})
	}
}
