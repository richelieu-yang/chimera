package confKit

import (
	"fmt"
	"testing"
	"time"
)

func TestMustLoad(t *testing.T) {
	type config struct {
		Time    time.Duration `json:"time"`
		Number  int           `json:"number,range=[1:100)"`
		Number1 int           `json:"number1,range=[1:]"`
	}

	c := &config{}
	MustLoad("./test.yaml", c)

	fmt.Printf("c.Time: %s\n", c.Time)
	fmt.Printf("c.Number: %d\n", c.Number)
	fmt.Printf("c.Number1: %d\n", c.Number1)
}
