package cmdKit

import (
	"context"
	"fmt"
	"testing"
)

func TestNewCommand(t *testing.T) {
	cmd := NewCommand(context.TODO(), "echo", nil)
	data, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func TestExecute(t *testing.T) {
	data, err := Execute("java", "-version")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Output:\n%s\n", string(data))
}
