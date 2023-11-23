package cmdKit

import (
	"context"
	"fmt"
	"testing"
)

func TestLookPath(t *testing.T) {
	path, err := LookPath("java")
	if err != nil {
		panic(err)
	}
	fmt.Println(path) // /usr/bin/java
}

func TestNewCommand(t *testing.T) {
	cmd := NewCommand(context.TODO(), "echo", nil)
	data, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func TestExecute(t *testing.T) {
	data, err := Execute(context.TODO(), "java", "-version")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Output:\n%s\n", string(data))
}
