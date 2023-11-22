package cmdKit

import (
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {
	data, err := Execute("java", "-version")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Output:\n%s\n", string(data))
}
